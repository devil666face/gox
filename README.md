# Gox

Gox is a lightweight proxy server with SOCKS5 and HTTPS (HTTP CONNECT) support, authentication, and optional systemd autostart. The project is written in Go, easily builds into a static binary, and can be run in Docker.

## Features

- SOCKS5 proxy with username/password authentication
- HTTPS (HTTP CONNECT) proxy with Basic authentication
- Simple configuration via YAML file or environment variables
- Self-signed TLS certificate generation for HTTPS proxy
- Optional autostart via systemd
- Ready-to-use Docker images and Docker Compose build

## Quick Start

> **To build the project, use the `task docker.build` command from the Taskfile.yml.**

### Build the binary

```bash
task build
```

Or build with obfuscation (garble):

```bash
task build.garble
```

### Generate certificates

```bash
task web.certs
```

### Run via Docker Compose

For testing:

```bash
docker compose -f docker-compose.test.yaml up --build
```

For production:

```bash
docker compose up -d
```

### Environment variables

To set the login and password, use environment variables during build:

- `BUILD_USERNAME` — user login
- `BUILD_PASSWORD` — user password

They will be embedded into the binary at build time via `-ldflags`.

### Example configuration file

```yaml
auth:
  username: password
socks:
  enable: true
  listen: "0.0.0.0:31080"
https:
  enable: true
  listen: "0.0.0.0:38443"
```

## Build and run via Docker

```bash
docker compose -f docker-compose.build.yaml up --build
```

After building, the binaries will be in the `bin/` folder.

To run the container, use:

```bash
docker compose up -d
```

## System dependencies

- Go 1.25+
- make, task, upx, objcopy, openssl (for building and packing)
- Docker (for containerization)

## Autostart via systemd

To set up autostart, run:

```bash
./gox -setup
```

To remove autostart:

```bash
./gox -remove
```

## Testing HTTPS proxy

```bash
curl -x https://127.0.0.1:38443 --proxy-user username:password --proxy-insecure -k https://ident.me
```

## Adding the certificate to trusted

Download the certificate via browser or use the generated `server.crt`:

```bash
mv cert.pem /usr/local/share/ca-certificates/gox.crt
update-ca-certificates
```

## Project structure

- `cmd/gox/` — entry point, main binary
- `internal/server/` — proxy implementation (SOCKS5, HTTPS), config, logic
- `pkg/` — helper packages (e.g., file operations, autostart)
- `Taskfile.yml` — build and generation tasks
- `Dockerfile`, `docker-compose.yaml` — containerization

## License

MIT

---

- https://github.com/lqqyt2423/go-mitmproxy
- https://github.com/elazarl/goproxy
- https://github.com/AdguardTeam/gomitmproxy
