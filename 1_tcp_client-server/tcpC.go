// во вторую очередь запускаем клиент go run tcpC.go 127.0.0.1:1234
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){

	arguments := os.Args

	if len(arguments)==1{
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err !=nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Print(">> ")
		text, err :=bufio.NewReader(os.Stdin).ReadString('\n')
		if err !=nil {
			fmt.Print(err)
		}

		_, _ = fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

