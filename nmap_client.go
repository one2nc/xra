package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Ullaakut/nmap"
)

var wgc sync.WaitGroup

func runNmapClient(hosts map[string]PortMap) {
	for h, p := range hosts {
		wg.Add(2)

		go func(h string, p []int, isUDP bool) {
			defer wg.Done()
			inner(makeScanner(h, p, isUDP))
		}(h, p["tcp"], false)

		go func(h string, p []int, isUDP bool) {
			defer wg.Done()
			inner(makeScanner(h, p, isUDP))
		}(h, p["udp"], true)
	}

	wg.Wait()
}

func makeScanner(host string, ports []int, isUDP bool) func(ctx context.Context) (*nmap.Scanner, error) {
	p := []string{}

	for _, i := range ports {
		p = append(p, strconv.Itoa(i))
	}

	return func(ctx context.Context) (*nmap.Scanner, error) {
		if isUDP {
			return nmap.NewScanner(
				nmap.WithTargets(host),
				nmap.WithPorts(strings.Join(p, ",")),
				nmap.WithUDPScan(),
				nmap.WithContext(ctx),
			)
		}
		return nmap.NewScanner(
			nmap.WithTargets(host),
			nmap.WithPorts(strings.Join(p, ",")),
			nmap.WithContext(ctx))

	}
}

func inner(f func(ctx context.Context) (*nmap.Scanner, error)) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	scanner, err := f(ctx)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Use the results to print an example output
	for _, h := range result.Hosts {
		if len(h.Ports) == 0 || len(h.Addresses) == 0 {
			continue
		}

		for _, port := range h.Ports {
			fmt.Printf(
				"%v, %v, %v -> %v:%v\n",
				port.State, port.Protocol, *host, h.Addresses[0], port.ID,
			)
		}
	}

	return nil
}
