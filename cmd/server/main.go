package main

import (
	"flag"
	"fmt"
	"log"

	"gox/internal/server/config"
	"gox/internal/server/gox"
	"gox/internal/server/socks"
)

var (
	configName = flag.String("config", "config.yaml", "path to config")
)

func main() {
	flag.Parse()

	_config, err := config.New(*configName)
	if err != nil {
		log.Fatalln("fatal config error:", err)
	}
	fmt.Println(_config)

	_socks := socks.New(
		_config.Proto,
		_config.Listen,
		_config.Auth,
	)

	_gox := gox.New(
		_socks,
	)

	if err := _gox.Listen(); err != nil {
		log.Fatalln("proxy fatal error:", err)
	}
}
