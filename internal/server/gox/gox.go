package gox

import "gox/internal/server/socks"

type Gox struct {
	server *socks.Socks
}

func New(
	_server *socks.Socks,
) *Gox {
	return &Gox{
		server: _server,
	}
}

func (g *Gox) Listen() error {
	return g.server.Listen()
}
