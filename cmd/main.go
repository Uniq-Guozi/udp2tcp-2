package main

import (
	"flag"
	"github.com/Error996/udp2tcp"
	"time"
	"flag"
)

var config struct {
	Client string
	PointA string
	Server string
	PointZ string
}

func main() {
	flag.StringVar(&config.Client, "c", "", "client listen address")
	flag.StringVar(&config.PointA, "a", "", "A point")
	flag.StringVar(&config.Server, "s", "", "server listen address")
	flag.StringVar(&config.PointZ, "z", "", "Z point")
	flag.Parse()

	if config.Client != "" {
		if config.PointA != "" {
			udp2tcp_client.StartClient(config.Client, config.PointA, time.Minute * 3)
		}
	}

	if config.Server != "" {
		if config.PointZ != "" {
			udp2tcp_server.StartServer(config.Server, config.PointZ, time.Minute * 3)
		}
	}
}
