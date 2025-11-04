package gox

import (
	"fmt"

	"gox/internal/server/config"
	"gox/internal/server/https"
	"gox/internal/server/socks"
)

type Gox struct {
	config *config.Config
	socks  *socks.Socks
	https  *https.Https
}

func New(
	_config *config.Config,
	key, cert []byte,
) (*Gox, error) {
	_gox := &Gox{
		config: _config,
	}
	if *_config.Socks.Enable {
		_gox.socks = socks.New(
			_config.Socks.Listen,
			_config.Auth,
		)
	}
	if *_config.Https.Enable {
		_https, err := https.New(
			_config.Https.Listen,
			_config.Auth,
			cert, key,
		)
		if err != nil {
			return nil, fmt.Errorf("https server fatal error: %w", err)
		}
		_gox.https = _https
	}
	return _gox, nil
}

func (g *Gox) Listen() error {
	var (
		i    int
		errs = make(chan error, 2)
	)

	if g.https != nil {
		i++
		go func() {
			errs <- g.https.Listen()
		}()
	}
	if g.socks != nil {
		i++
		go func() {
			errs <- g.socks.Listen()
		}()
	}
	if i == 0 {
		return fmt.Errorf("no servers enabled")
	}
	return <-errs
}
