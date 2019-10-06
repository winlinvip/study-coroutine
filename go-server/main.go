package main

import (
	"fmt"
	"net"
)

const ListenPort = 8080

func serve(conn *net.TCPConn) {
	fmt.Println("start serve client")
	buf := make([]byte, 1024)
	var nnMsgs uint64
	for {
		nn, err := conn.Read(buf)
		if err != nil {
			break
		}

		_, err = conn.Write(buf[:nn])
		if err != nil {
			break
		}

		nnMsgs++
	}

	fmt.Println(fmt.Sprintf("server done, msgs=%v", nnMsgs))
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%v", ListenPort))
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Listen=%v", ListenPort))

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go serve(conn)
	}
}
