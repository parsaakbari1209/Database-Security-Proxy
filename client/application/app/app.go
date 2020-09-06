package app

import (
	"context"
	"io"
	"log"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb"
	"google.golang.org/grpc"
)

var (
	c tlspb.TLSClientServiceClient
)

func StartService() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect client-side-tls-service: %v", err)
	}
	defer conn.Close()
	c = tlspb.NewTLSClientServiceClient(conn)
}

func SendReq(connString, query string) (bool, string, string) {
	// Request that is going to be sent to client-side-tls-service.
	req := &tlspb.TLSClientRequest{
		ConnString: connString,
		SqlString:  query,
	}
	// Calling the releated RPC with no deadlines.
	res, err := c.TLSClientSend(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling TLSClientSend RPC: %v", err)
	}
	// Reciving streams and
	for {
		msg, err := res.Recv()
		if err == io.EOF { // End of streaming.
			break
		}
		if err != nil { // Other errors.
			log.Fatalf("error while server-side-tls-service streaming: %v", err)
		}
		return msg.GetSucceed(), msg.GetResult(), msg.GetError()
	}
	// Notify that streaming is finished.
	return false, "", "EOF"
}
