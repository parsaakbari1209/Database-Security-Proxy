package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	}
}
