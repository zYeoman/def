package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func def(conn net.Conn, s string) {
	i := 0
	buffer := make([]byte, 1024)
	tosend := []byte(s)
	size := len(tosend)
	size_send := make([]byte, 1)
	size_send[0] = byte(size)
	conn.Write(size_send)
	conn.Write(tosend)
	for {
		n, err := conn.Read(buffer)
		fmt.Println(string(buffer[:n]))
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error: ", err)
			}
			break
		}
		if i >= 10 {
			break
		}
		i++
	}
}

func sender(conn net.Conn) {
	defer conn.Close()
	words := os.Args[1]
	def(conn, words)
}

func main() {
	server := "127.0.0.1:3154"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	sender(conn)

}
