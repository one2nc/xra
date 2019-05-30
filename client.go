package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func runClient(hosts map[string]PortMap) {
	for h, p := range hosts {
		for proto, ports := range p {
			for _, port := range ports {
				wg.Add(1)
				if proto == "tcp" {
					go func(proto, h string, port int) {
						defer wg.Done()
						if err := clientTCPConnect(h, port); err != nil {
							fmt.Printf("fail, %v, %v -> %v:%v\n", proto, *host, h, port)
							return
						}

						fmt.Printf("pass, %v, %v -> %v:%v\n", proto, *host, h, port)
					}(proto, h, port)
				} else {
					go func(proto, h string, port int) {
						defer wg.Done()
						if err := clientUDPConnect(h, port); err != nil {
							fmt.Printf("fail, %v, %v -> %v:%v\n", proto, *host, h, port)
							return
						}
						fmt.Printf("pass, %v, %v -> %v:%v\n", proto, *host, h, port)

						//log.Println("Success:", h, proto, port)
					}(proto, h, port)
				}
			}
		}
	}
	wg.Wait()
}

func clientTCPConnect(host string, port int) error {
	// log.Printf("connecting to %v://%v:%v\n", protocol, host, port)
	con, err := net.DialTimeout(
		"tcp",
		fmt.Sprintf("%v:%v", host, port),
		5*time.Second,
	)

	if err != nil {
		return err
	}

	defer con.Close()
	return nil
}

func clientUDPConnect(host string, port int) error {
	protocol := "udp"

	//log.Printf("connecting to %v://%v:%v\n", protocol, host, port)
	conn, err := net.Dial(protocol, fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return err
	}

	go func(conn net.Conn) {
		<-time.After(5 * time.Second)
		conn.Close()
	}(conn)

	defer conn.Close()
	fmt.Fprintf(conn, fmt.Sprintf("UDP Request sent to %v", port))
	buffer := make([]byte, 1024)
	if _, err := conn.Read(buffer); err != nil {
		return err
	}
	return nil
}
