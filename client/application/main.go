package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/parsaakbari1209/Database-Security-Proxy/client/application/app"
)

// Main func is a shell like programme running in a terminal.
// It gets a connection string and a query and sends them to SendReq func in app/app.go
// Then gets the result and prints in them.
// These steps are looped forever until the proccess is killed.
func main() {
	fmt.Println("started client command-line application...")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("con str -> ")
		connString, _ := reader.ReadString('\n')
		connString = strings.Replace(connString, "\n", "", -1)

		fmt.Print("sql str -> ")
		query, _ := reader.ReadString('\n')
		query = strings.Replace(query, "\n", "", -1)

		fmt.Println("")
		fmt.Println("---------- RESULT ----------")
		res := app.SendReq(connString, query)
		for _, row := range res {
			fmt.Println(row[1])
		}
		fmt.Println("")
	}
}
