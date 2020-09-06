package app

import (
	"context"
	"io"
	"log"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb"
)

// SendReq func sends a gRPC request to the client-side-tls-service.
// Then, returns the stream that is coming from client-side-tls-service as a 2D array.
func SendReq(connString, query string, c tlspb.TLSClientServiceClient) [][]string {
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
	// TODO: Add streaming functionality for responding the function call.
	// Recive stream and, return the results as a 2D array.
	result := [][]string{}
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
		// Append stream result to the 2D array.
		result = append(result, []string{
			boolToString(msg.GetSucceed()),
			msg.GetResult(),
			msg.GetError(),
		})
	}
	// Return the 2D array.
	return result
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
