package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/app"
	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb"
	"google.golang.org/grpc"
)

var (
	// c (package-level variable) is the connection to the client-side-tls-service.
	c tlspb.TLSClientServiceClient
)

// main func is a shell like programme running in a terminal.
// First, connects to the client-side-tls-service server.
// Second, gets a connection string and a query.
// Third, calls SendReq func with connection-string, query and connection-iteself.
// Fourth, gets the result(stream) and prints them to standard-output.
// These steps are looped forever until the proccess is killed.
func main() {
	fmt.Println("started client command-line application...")
	fmt.Println("")

	// Reader for reading lines from terminal.
	reader := bufio.NewReader(os.Stdin)

	// Dial to the client-side-tls-service server.
	address := "localhost:50051"
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect client-side-tls-service: %v", err)
	}
	defer cc.Close()
	c = tlspb.NewTLSClientServiceClient(cc)

	// Get newline char based on the OS.
	newline := newline()
	// Loop forever for shell like functionality.
	for {
		// Get connection string.
		fmt.Print("con str -> ")
		connString, _ := reader.ReadString('\n')
		connString = strings.Replace(connString, newline, "", -1)

		// Get query.
		fmt.Print("sql str -> ")
		query, _ := reader.ReadString('\n')
		query = strings.Replace(query, newline, "", -1)

		// Send request and show response.
		fmt.Println("")
		fmt.Println("---------- RESULT ----------")
		r := app.SendReq(connString, query, c)
		if _, err := io.Copy(os.Stdout, r); err != nil {
			log.Fatalf("error while writing to standard-output: %v", r)
		}
		fmt.Println("")
	}
}

// newline func returns a newline character based on the os.
// For windows OSes: \r\n
// For unix-like OSes: \n
func newline() string {
	newline := "\n"
	if runtime.GOOS == "windows" {
		newline = "\r\n"
	}
	return newline
}
