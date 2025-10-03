package request

import (
	"github.com/things-go/go-socks5"
)

const username = "username"

type Request struct {
	Username string
}

func New(req *socks5.Request) *Request {
	return &Request{
		Username: req.AuthContext.Payload[username],
	}
}
