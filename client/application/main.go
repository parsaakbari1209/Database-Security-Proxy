package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/app"
	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb"
	"google.golang.org/grpc"
)

var (
	// c (package-level variable) is the connection to the client-side-tls-service.
	c tlspb.TLSClientServiceClient
)

// TODO: the stream functionality must be implemented through a pipe!
// main func is a shell like programme running in a terminal.
// First, connects to the client-side-tls-service server.
// Second, gets a connection string and a query and sends them to SendReq func with c.
// Third, gets the result and prints them.
// These steps are looped forever until the proccess is killed.
func main() {
	fmt.Println("started client command-line application...")
	fmt.Println("")

	// Reader for reading lines from terminal.
	reader := bufio.NewReader(os.Stdin)

	// TODO: Implemetent secure connection!
	// Dial to the client-side-tls-service server.
	address := "localhost:50051"
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect client-side-tls-service: %v", err)
	}
	defer cc.Close()
	c = tlspb.NewTLSClientServiceClient(cc)

	// Loop forever for shell like functionality.
	for {
		// TODO: Newlines in windows are \r\n so, this won't work on windows.
		// Get connection string.
		fmt.Print("con str -> ")
		connString, _ := reader.ReadString('\n')
		connString = strings.Replace(connString, "\n", "", -1)

		// TODO: Newlines in windows are \r\n so, this won't work on windows.
		// Get query.
		fmt.Print("sql str -> ")
		query, _ := reader.ReadString('\n')
		query = strings.Replace(query, "\n", "", -1)

		// Send request and show response.
		fmt.Println("")
		fmt.Println("---------- RESULT ----------")
		res := app.SendReq(connString, query, c)
		for _, row := range res {
			fmt.Println(row[1])
		}
		fmt.Println("")
	}
}
