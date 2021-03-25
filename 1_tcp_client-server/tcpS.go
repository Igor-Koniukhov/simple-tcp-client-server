// запускаем в первую очередь сервер  go run tcpS.go 1234

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)


func main(){
	arguments := os.Args
	fmt.Println(arguments)
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	checkErrRet(err)
	defer l.Close()

	c, err := l.Accept()
	checkErrRet(err)

	checkErrRet(err)
	fl, err := os.OpenFile("info.txt", os.O_RDWR, 0666,)
	checkErrRet(err)

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		checkErrRet(err)

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("-> ", string(netData))

		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))

		 fl.Write([]byte(netData + myTime))

	}

}

func checkErrRet(err error){
	if err !=nil {
		fmt.Println(err)
		return
	}
}