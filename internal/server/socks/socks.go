package socks

import (
	"gox/internal/server/socks/ruler"
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

type Socks struct {
	server *socks5.Server
	proto  string
	listen string
}

func New(
	_proto string,
	_listen string,
	creds map[string]string,
) *Socks {
	return &Socks{
		proto:  _proto,
		listen: _listen,
		server: socks5.NewServer(
			socks5.WithAuthMethods(
				[]socks5.Authenticator{
					socks5.UserPassAuthenticator{
						Credentials: toCreds(creds),
					},
				},
			),
			socks5.WithLogger(
				socks5.NewLogger(
					log.New(os.Stdout, "", log.LstdFlags),
				),
			),
			socks5.WithRule(ruler.New()),
		),
	}
}

func (s *Socks) Listen() error {
	log.Printf("listen on %s/%s", s.listen, s.proto)
	return s.server.ListenAndServe(
		s.proto,
		s.listen,
	)
}

func toCreds(creds map[string]string) socks5.StaticCredentials {
	var staticCreds = make(socks5.StaticCredentials)
	for username, password := range creds {
		staticCreds[username] = password
	}
	return staticCreds
}
