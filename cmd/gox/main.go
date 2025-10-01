package main

import (
	"flag"
	"log"

	_ "embed"

	"gox/internal/server/config"
	"gox/internal/server/gox"
	"gox/pkg/start"
)

var (
	path   = flag.String("config", "config.yaml", "path to config")
	save   = flag.Bool("save", false, "save default config to path -config")
	setup  = flag.Bool("setup", false, "set autostart via systemd")
	remove = flag.Bool("remove", false, "remove autostart via systemd")
)

//go:embed server.key
var key []byte

//go:embed server.crt
var cert []byte

var (
	Username, Password string
)

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

	if _start, err := start.New(); err == nil {
		if *setup {
			if err := _start.Setup(); err != nil {
				log.Fatalln("set autostart error:", err)
			}
			return
		}
		if *remove {
			if err := _start.Remove(); err != nil {
				log.Fatalln("remove autostart error:", err)
			}
			return
		}
	} else {
		log.Println("init autostart error:", err)
	}

	_config, err := config.New(
		*path,
		Username,
		Password,
	)
	if err != nil {
		log.Fatalln("fatal config error:", err)
	}

	_gox, err := gox.New(
		_config,
		key, cert,
	)
	if err != nil {
		log.Fatalln("create fatal error:", err)
	}

	if err := _gox.Listen(); err != nil {
		log.Fatalln("proxy fatal error:", err)
	}
}
