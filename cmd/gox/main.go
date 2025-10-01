package main

import (
	"flag"
	"log"

	_ "embed"

	"gox/internal/server/config"
	"gox/internal/server/gox"
)

var (
	path = flag.String("config", "config.yaml", "path to config")
	save = flag.Bool("save", false, "save default config to path -config")
)

//go:embed server.key
var Key []byte

//go:embed server.crt
var Cert []byte

func main() {
	flag.Parse()

	if *save {
		if err := config.Default(
			*path,
		); err != nil {
			log.Fatalln(err)
		}
		return
	}

	_config, err := config.New(*path)
	if err != nil {
		log.Fatalln("fatal config error:", err)
	}

	_gox, err := gox.New(
		_config,
		Key, Cert,
	)
	if err != nil {
		log.Fatalln("create fatal error:", err)
	}

	if err := _gox.Listen(); err != nil {
		log.Fatalln("proxy fatal error:", err)
	}
}
