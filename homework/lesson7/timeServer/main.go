package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	go exit()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func exit() {

	for {
		var exitFlag string
		fmt.Scan(&exitFlag)
		if exitFlag == "exit" {
			os.Exit(0)
		} else {
			fmt.Println("maybe you whant type exit")
		}

	}
}
