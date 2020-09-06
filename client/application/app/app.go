package app

import (
	"context"
	"io"
	"log"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb"
	"google.golang.org/grpc"
)

var (
	// c package-level variable is the connection to client-side-tls-service.
	c tlspb.TLSClientServiceClient
)

// StartService func dials to the client-side-tls-service.
func StartService() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect client-side-tls-service: %v", err)
	}
	defer conn.Close()
	c = tlspb.NewTLSClientServiceClient(conn)
}

// SendReq func sends a gRPC request to the client-side-tls-service.
// Then returns the stream that is coming from client-side-tls-service as an 2D array.
func SendReq(connString, query string) [][]string {
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
	// Recive stream and return the results as array.
	result := [][]string{}
	for {
		msg, err := res.Recv()
		if err == io.EOF { // End of streaming.
			break
		}
		if err != nil { // Other errors.
			log.Fatalf("error while server-side-tls-service streaming: %v", err)
		}
		if msg.GetSucceed() == false { // Check if the operation was successful.
			log.Fatalf("error is occured and stream is not succeed: %v", msg.GetError())
		}
		result = append(result, []string{
			boolToString(msg.GetSucceed()),
			msg.GetResult(),
			msg.GetError(),
		})
	}
	return result
}

// boolToString func returns the bool as a string value.
func boolToString(theBool bool) string {
	res := "false"
	if theBool == true {
		res = "true"
	}
	return res
}
