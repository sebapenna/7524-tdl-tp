# 7524 - Teoría del Lenguaje: Go
- Cleanup unused dependencies: `go mod tidy`

## Installation Guide
- On _root_ folder of the project run the following commands
```bash
$ go build ./cmd/...
$ go install ./cmd/...
```

## Server
- Accepts multiple connections

### Run
- App installed: `fiubados-server <port_number>`
- App not installed: `go run cmd/fiubados-server/server.go <port_number>`

### Exit
- Type `EXIT` on prompt

## Client

### Run
- App installed: `fiubados-client <ip:port_number>`
- App not installed: `go run cmd/fiubados-client/client.go <ip:port_number>`

### Exit
- Type `STOP` on prompt
