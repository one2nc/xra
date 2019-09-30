package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func hostname() string {
	h, e := os.Hostname()
	if e != nil {
		panic(e)
	}
	return h
}

var (
	debug    = flag.Bool("debug", false, "Run in debug mode?")
	use_nmap = flag.Bool("nmap", false, "Nmap scan? [Required sudo]")
	client   = flag.Bool("client", false, "Run in client mode?")
	config   = flag.String("config", "/tmp/goss.json", "Path to Json")
	host     = flag.String("host", hostname(), "Hostname or IP")
)

type PortConf struct {
	AllowTCP       []int `json:"allow_tcp"`
	AllowUDP       []int `json:"allow_udp"`
	ThroughputTest []int `json:"throughputTest"`
}

type MasterConfig struct {
	Zones   map[string][]string            `json:"zones"`
	Network map[string]map[string]PortConf `json:"network"`
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Printf("\nHere")
	var mConfig MasterConfig
	mb, err := ioutil.ReadFile(*config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", mb)
	if err := json.Unmarshal(mb, &mConfig); err != nil {
		panic(err)
	}

	if *client {
		c := spitClientConfig(&mConfig, *host)
		fmt.Printf("\nIn Client mode, config is %v", c)
		if *debug {
			log.Printf("%+v\n", c)
		}
		if *use_nmap {
			runNmapClient(c)
		} else {
			runClient(c)
		}
		return
	}

	c := spitServerConfig(&mConfig, *host)
	fmt.Printf("\nServer config%+v\n", c)
	if *debug {
		log.Printf("%+v\n", c)
	}
	runServer(c)
}
