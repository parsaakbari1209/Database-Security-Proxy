package app

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb"
)

// SendReq func sends a gRPC request to the client-side-tls-service.
// Then, returns a pipe reader for streaming back.
func SendReq(connString, query string, c tlspb.TLSClientServiceClient) *io.PipeReader {
	// Request that is going to be sent to client-side-tls-service.
	req := &tlspb.TLSClientRequest{
		ConnString: connString,
		SqlString:  query,
	}

	// Call the releated RPC with no deadlines.
	res, err := c.TLSClientSend(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling TLSClientSend RPC: %v", err)
	}

	// Create a pipe for streaming back.
	pr, pw := io.Pipe()

	// Start to recive streams from client-side-tls-service in a new goroutine.
	// And write them to the pipe.
	go func() {
		// Loop until end of file.
		for {
			msg, err := res.Recv()
			if err == io.EOF {
				// End of streaming.
				break
			}
			if err != nil {
				// Other errors.
				log.Fatalf("error while server-side-tls-service streaming: %v", err)
			}
			// Check if the operation was successful.
			if msg.GetSucceed() == false {
				log.Fatalf("error is occured and stream is not succeed: %v", msg.GetError())
			}
			// Write the stream result to the pipe.
			fmt.Fprint(pw, msg.GetResult()+"\n")
		}
		pw.Close()
	}()
	// Return pipe reader.
	return pr
}

// boolToString func returns the bool as a string value.
// "false" for false boolean.
// "true" for true boolean.
func boolToString(theBool bool) string {
	res := "false"
	if theBool == true {
		res = "true"
	}
	return res
}
