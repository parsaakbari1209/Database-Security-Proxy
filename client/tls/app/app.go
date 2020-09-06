package app

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/tls/tlspb"
	"google.golang.org/grpc"
)

type server struct{}

// TODO: Not implemented (It is a mock.).
// Send Client request from client-side-tls-service to server-side-tls-service and return response.
func (s *server) TLSClientSend(req *tlspb.TLSClientRequest, stream tlspb.TLSClientService_TLSClientSendServer) error {
	fmt.Println("TLSClientSend RPC invoked.")
	// Client Request data extraction.
	//connString := req.GetConnString()
	//sqlString := req.GetSqlString()
	// Call server-side-tls-service with a new request.
	// Not implemented...
	// Return server-side-tls-service stream to the client.
	// Not implemented...

	// Mock anwser.
	for i := 0; i < 10; i++ {
		res := &tlspb.TLSClientResponse{
			Succeed: true,
			Result:  "Number: " + strconv.Itoa(i),
			Error:   "",
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
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
