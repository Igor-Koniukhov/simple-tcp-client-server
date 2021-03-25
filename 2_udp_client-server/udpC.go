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
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]
	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	checkErr(err)
	c, err := net.DialUDP("udp4", nil, s)
	checkErr(err)

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()
	for {
		fmt.Print(">> ")

		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		data := []byte(text + "\n")
		_, err = c.Write(data)
		/*fmt.Fprintf(c, text + "\n")*/

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}
		checkErr(err)

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		checkErr(err)

		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}

}


func checkErr(err error){
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
}