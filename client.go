package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func clientThrougputTest(host string, port int) error {
	conn, err := net.DialTimeout(
		"tcp",
		fmt.Sprintf("%v:%v", host, port),
		5*time.Second,
	)

	if err != nil {
		fmt.Fprintf(os.Stdout, "\nError while opening connection, error is %s", err)
		return err
	}

	message, err := createRandomMessage(PERF_MESSAGE_LENGTH)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(-1)
	}

	enc := gob.NewEncoder(conn)
	messageHeader := GetPerfMessageHeader(JUST_HELLO, 0)
	enc.Encode(messageHeader)

	fmt.Printf("\nSending PERF_START")
	messageHeader = GetPerfMessageHeader(PERF_START, 0)
	enc.Encode(messageHeader)

	// Send the payload messages
	for i := 0; i < 100; i++ {
		fmt.Printf("\nSending PERF_PAYLOAD %d", i)
		messageHeader = GetPerfMessageHeader(PERF_PAYLOAD, PERF_MESSAGE_LENGTH)
		enc.Encode(messageHeader)

		/*
		 * conn.Write(message)
		 */
		enc.Encode(message)
	}

	fmt.Printf("\nSending PERF_END")
	messageHeader = GetPerfMessageHeader(PERF_END, 0)
	enc.Encode(messageHeader)

	fmt.Printf("\nSending JUST_BYE")
	messageHeader = GetPerfMessageHeader(JUST_BYE, 0)
	enc.Encode(messageHeader)
	conn.Close()

	/*
	 * time.Sleep(10 * time.Second)
	 */
	return nil
}

func runClient(hosts map[string]PortMap) {
	fmt.Printf("\nIn Run client\n")
	fmt.Printf("Hosts %+v", hosts)

	for h, p := range hosts {
		for proto, ports := range p {
			for _, port := range ports {
				fmt.Printf("\nProto is %s", proto)
				wg.Add(1)
				if proto == "throughputTest" {
					fmt.Printf("\nIn throughput test")
					clientThrougputTest(h, port)
				}
				if proto == "tcp" {
					go func(proto, h string, port int) {
						defer wg.Done()
						fmt.Printf("\nConnecting to host %s port %d on TCP", h, port)
						if err := clientTCPConnect(h, port); err != nil {
							fmt.Printf("fail, %v, %v -> %v:%v\n", proto, *host, h, port)
							return
						}

						fmt.Printf("pass, %v, %v -> %v:%v\n", proto, *host, h, port)
					}(proto, h, port)
				} else {
					go func(proto, h string, port int) {
						defer wg.Done()
						fmt.Printf("\nConnecting to host %s port %d on UDP", h, port)
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
