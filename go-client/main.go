package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	var nnClients,interval int
	flag.IntVar(&nnClients, "c", 1, "The number of clients")
	flag.IntVar(&interval, "i", 60, "The interval in seconds")
	flag.Parse()
	fmt.Println(fmt.Sprintf("clients=%v, interval=%vs", nnClients, interval))

	var wg sync.WaitGroup
	for i := 0; i < nnClients; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Int()%5000) * time.Millisecond)

			addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
			if err != nil {
				fmt.Println("err is", err)
				return
			}

			conn, err := net.DialTCP("tcp", nil, addr)
			if err != nil {
				fmt.Println("err is", err)
				return
			}

			buf := make([]byte, 1024)
			for {
				if _, err := conn.Write([]byte("Ping")); err != nil {
					fmt.Println("err is", err)
					return
				}

				_, err := conn.Read(buf)
				if err != nil {
					fmt.Println("err is", err)
					return
				}

				time.Sleep(time.Duration(interval) * time.Second)
			}
		}()
	}

	wg.Wait()
}
