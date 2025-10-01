package https

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

type Https struct {
	server *goproxy.ProxyHttpServer
	listen string
	creds  map[string]string
	tls    *tls.Config
}

func New(
	_listen string,
	_creds map[string]string,
	cert, key []byte,
) (*Https, error) {
	_tls, err := loadTLSCreds(cert, key)
	if err != nil {
		return nil, fmt.Errorf("failed to load tls: %w", err)
	}
	return &Https{
		server: goproxy.NewProxyHttpServer(),
		listen: _listen,
		creds:  _creds,
		tls:    _tls,
	}, nil
}

func authMiddleware(auth map[string]string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Proxy-Authorization")
		if header == "" {
			w.Header().Set("Proxy-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
			return
		}
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			w.Header().Set("Proxy-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			w.Header().Set("Proxy-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
			return
		}
		creds := strings.SplitN(string(decoded), ":", 2)
		if len(creds) != 2 {
			w.Header().Set("Proxy-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
			return
		}
		username, password := creds[0], creds[1]
		pass, ok := auth[username]
		if !ok || pass != password {
			w.Header().Set("Proxy-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Proxy Authentication Required", http.StatusProxyAuthRequired)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Https) Listen() error {
	log.Printf("listen https on %s", s.listen)
	server := &http.Server{
		Addr:      s.listen,
		TLSConfig: s.tls,
		Handler:   authMiddleware(s.creds, s.server),
	}
	listen, err := tls.Listen("tcp", s.listen, s.tls)
	if err != nil {
		return fmt.Errorf("failed to start TLS listener: %w", err)
	}
	return server.Serve(listen)
}

func loadTLSCreds(cert, key []byte) (*tls.Config, error) {
	creds, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, fmt.Errorf("failed to load server cert and key: %w", err)
	}
	return &tls.Config{
		Certificates:       []tls.Certificate{creds},
		InsecureSkipVerify: true,
	}, nil
}
