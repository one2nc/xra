package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type PortMap map[string][]int

type clientConfig struct {
	throughputTest bool
	portMap        map[string]PortMap
}

func sendTcpResponse(conn net.Conn) {
	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	conn.Write([]byte("Message received."))
	conn.Close()
}

func tcpListen(port int) {
	fmt.Printf("\ntcpListenListening on the port %d", port)
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Println(err)
		return
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}
		go sendTcpResponse(conn)
	}
}

func serverThroughputTest(port int) {
	fmt.Printf("\nListening on the port %d", port)
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		fmt.Printf("\nError Listening")
		log.Println(err)
		return
	}

	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		log.Println("Error accepting: ", err.Error())
	}
	__serverThroughputTest(conn)
}

func __serverThroughputTest(conn net.Conn) {

	fmt.Fprintf(os.Stderr, "\nCalling __serverThroughputTest")

	var data PerfMessageHeader
	var counter int

	decoder := gob.NewDecoder(conn)
	var startTimeStamp time.Time
	var transferredBytes int
	for {
		err := decoder.Decode(&data)

		if err != nil {
			fmt.Printf("\nError while decoding before switch, test failed, return, %s", err)
			return
		}

		/*
		 * fmt.Printf("\nRecieved data %+v\n", data)
		 */
		switch data.Type {
		case JUST_HELLO:
		case PERF_START:
			fmt.Printf("\nStarting performance testing")
			startTimeStamp = data.TimeStamp
		case PERF_PAYLOAD:
			buffer := make([]byte, data.Length)

			err := decoder.Decode(&buffer)

			if err != nil {
				fmt.Printf("\nPERF_PAYLOAD Error while decoding, test failed, return in PERF_PAYLOAD, %s", err)
				return
			}

			/*
			 * readlen, err := io.ReadAtLeast(conn, buffer, data.Length)
			 * _, _ = readlen, err
			 */

			transferredBytes += data.Length

			counter++

			if counter%10 == 0 {
				fmt.Printf("\nGot %d messages", counter)
			}

		case PERF_END:
			dur := time.Since(startTimeStamp)
			fmt.Printf("\nGOT PERF END, will print the summary")
			fmt.Printf("\nTransferred %d bytes in %+v seconds", data.Length, dur.Seconds())
			transferredBytesMB := transferredBytes / (1024 * 1024 * 1024)
			fmt.Printf("\nSummary %d MB in %+v seconds Throughput %f MBPS", transferredBytesMB, dur.Seconds(), float64(transferredBytesMB)/dur.Seconds())
		case JUST_BYE:
			fmt.Printf("\nClosing connection, got JUST_BYE")
			conn.Close()
			return
		}
	}
}
func udpListen(port int) {
	// listen to incoming udp packets
	pc, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		log.Println(err)
		return
	}

	defer pc.Close()
	for {
		//simple read
		buffer := make([]byte, 1024)
		_, addr, err := pc.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Received:", string(buffer))

		//simple write
		pc.WriteTo([]byte("Hello from client"), addr)
	}
}

func server(p *PortMap) {
	for proto, ports := range *p {
		for _, port := range ports {
			log.Printf("Trying to listen on %v Port: %v", proto, port)
			if proto == "tcp" {
				/*
				 * tcpListen(port)
				 */
			} else if proto == "throughputTest" {
				serverThroughputTest(port)
			} else {
				/*
				 * udpListen(port)
				 */
			}
		}
	}
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()
	<-done
	log.Println("Exiting")
}

func runServer(x *PortMap) {
	server(x)
	waitForSignal()
}
