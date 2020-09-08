package app

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/parsaakbari1209/Database-Security-Proxy/server/tls/tlspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

// TODO: Not implemented (It is a mock.).
// TLSServerSend gets client-side-tls-request.
// First, decrypes the request.
// Second, sends the request to database-security-service.
// Third, gets stream response from database-security-service.
// Forth, encryptes the reponses and streams to client-side-tls-service.
func (s *server) TLSServerSend(req *tlspb.TLSServerRequest, stream tlspb.TLSServerService_TLSServerSendServer) error {
	fmt.Println("TLSServerSend RPC invoked.")
	// Mock code!
	for i := 0; i < 10; i++ {
		res := &tlspb.TLSServerResponse{
			Succeed: true,
			Result:  "Number: " + strconv.Itoa(i),
			Error:   "",
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	return nil
}

// TODO: Address and port number must exported from env vars or terminal args.
// StartService func start to listen on a port.
// The transport protocol is tcp.
func StartService() {
	fmt.Println("started server-side-tls-service...")

	// Configure a listener on port 50052.
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Load server certification and private key file and create credentials.
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("error while loading certificate or key file: %v", sslErr)
	}
	// Create options for new server from credentials.
	opts := grpc.Creds(creds)

	// Create a new grpc server that serves server-side-tls-service.
	s := grpc.NewServer(opts)
	tlspb.RegisterTLSServerServiceServer(s, &server{})

	// Start serving.
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
