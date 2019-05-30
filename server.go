package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type PortMap map[string][]int

func sendTcpResponse(conn net.Conn) {
	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	conn.Write([]byte("Message received."))
	conn.Close()
}

func tcpListen(port int) {
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
				go tcpListen(port)
			} else {
				go udpListen(port)
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
