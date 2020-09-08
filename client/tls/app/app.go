package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/tls/tlspb"
	"google.golang.org/grpc"
)

// server struct implementes the TLSClientServiceServer interface.
type server struct{}

// First, send client-request from client-side-tls-service to server-side-tls-service.
// Second, return response(stream back).
func (s *server) TLSClientSend(req *tlspb.TLSClientRequest, stream tlspb.TLSClientService_TLSClientSendServer) error {
	fmt.Println("TLSClientSend RPC invoked.")

	// Extract data from client-request.
	connString := req.GetConnString()
	sqlString := req.GetSqlString()

	// TODO: tls not implemented!
	// Dial to the server-side-tls-service server.
	address := "localhost:50052"
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect server-side-tls-service: %v", err)
	}
	defer cc.Close()
	c := tlspb.NewTLSServerServiceClient(cc)

	// Create request for the TLSServerSend RPC.
	newReq := &tlspb.TLSServerRequest{
		ConnString: connString,
		SqlString:  sqlString,
	}

	// Call server-side-tls-service TLSServerSend RPC.
	res, err := c.TLSServerSend(context.Background(), newReq)
	if err != nil {
		log.Fatalf("error while calling TLSServerSend RPC: %v", err)
	}

	// Return server-side-tls-service stream back to the client.
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			// End of stream.
			break
		}
		if err != nil {
			// Other errors.
			log.Fatalf("error while reciving stream: %v", err)
		}
		// Create new responses to send the client.
		newRes := &tlspb.TLSClientResponse{
			Succeed: msg.GetSucceed(),
			Result:  msg.GetResult(),
			Error:   msg.GetError(),
		}
		stream.Send(newRes)
	}
	return nil
}

// TODO: Address and port number must be exported from env vars or terminal args.
// StartService func starts to listen on a port.
// The transport protocol is tcp.
func StartService() {
	fmt.Println("started client-side-tls-service...")

	// Configure a listener on port 50051.
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new grpc server that serves client-side-tls-service.
	s := grpc.NewServer()
	tlspb.RegisterTLSClientServiceServer(s, &server{})

	// Start serving.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
