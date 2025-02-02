package request

import (
	"log"

	"github.com/things-go/go-socks5"
)

const username = "username"

type Request struct {
	Username string
}

func New(req *socks5.Request) *Request {
	log.Println(req.DstAddr)
	log.Println(req.LocalAddr)
	log.Println(req.RemoteAddr)
	log.Println(req.DestAddr)
	return &Request{
		Username: req.AuthContext.Payload[username],
	}
}
