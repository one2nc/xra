package main

import (
	"encoding/json"
	"flag"
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
	client = flag.Bool("client", false, "Run in client mode?")
	config = flag.String("config", "/tmp/goss.json", "Path to Json")
	host   = flag.String("host", hostname(), "Hostname or IP")
)

type MasterConfig struct {
	Zones   map[string][]string                      `json:"zones"`
	Network map[string]map[string]map[string][]int64 `json:"network"`
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var mConfig MasterConfig
	mb, err := ioutil.ReadFile(*config)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(mb, &mConfig); err != nil {
		panic(err)
	}

	if *client {
		runClient(spitClientConfig(&mConfig, *host))
		return
	}

	runServer(spitServerConfig(&mConfig, *host))
}
