package socks

import (
	"log"
	"os"

	"gox/internal/server/socks/ruler"

	"github.com/things-go/go-socks5"
)

type Socks struct {
	server *socks5.Server
	listen string
}

func New(
	_listen string,
	creds map[string]string,
) *Socks {
	return &Socks{
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
			socks5.WithResolver(socks5.DNSResolver{}),
		),
	}
}

func (s *Socks) Listen() error {
	log.Printf("listen socks on %s", s.listen)
	return s.server.ListenAndServe(
		"tcp",
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

// type SystemResolver struct{}

// func (r *SystemResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
// 	ips, err := net.DefaultResolver.LookupIP(ctx, "ip", name)
// 	if err != nil {
// 		return ctx, nil, err
// 	}
// 	return ctx, ips[0], nil
// }

// func (r *SystemResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
// 	resolver := &net.Resolver{
// 		PreferGo: true,
// 		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
// 			return net.Dial("udp", "127.0.0.11:53")
// 		},
// 	}
// 	ips, err := resolver.LookupIP(ctx, "ip", name)
// 	if err != nil {
// 		return ctx, nil, err
// 	}
// 	return ctx, ips[0], nil
// }
