package app

import (
	"fmt"
	"log"
	"net"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/tls/tlspb"
	"google.golang.org/grpc"
)

type server struct{}

// Send Client request from client-side-tls-service to server-side-tls-service and return response.
func (s *server) TLSClientSend(req *tlspb.TLSClientRequest, stream tlspb.TLSClientService_TLSClientSendServer) error {
	// Client Request data extraction.
	//connString := req.GetConnString()
	//sqlString := req.GetSqlString()
	// Call server-side-tls-service with a new request.
	// Not implemented...
	// Return server-side-tls-service stream to the client.
	// Not implemented...
	return nil
}

func StartService() {
	fmt.Println("started client-side-tls-service...")
	// Configure a listener on port 50051.
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	// Create a new grpc server that serves client-side-tls-service.
	s := grpc.NewServer()
	tlspb.RegisterTLSClientServiceServer(s, &server{})

	// Start serving.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
