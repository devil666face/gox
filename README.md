### Test https

curl -x https://127.0.0.1:8443 --proxy-user username:password --proxy-insecure -k https://ident.me

### Add certs

Download cert via browser

```bash
mv cert.pem /usr/local/share/ca-certificates/gox.crt
update-ca-certificates
```

---

- https://github.com/lqqyt2423/go-mitmproxy
- https://github.com/elazarl/goproxy
- https://github.com/AdguardTeam/gomitmproxy
