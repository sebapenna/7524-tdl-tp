# 7524 - Teoría del Lenguaje: Go

- Cleanup unused dependencies: `go mod tidy`

## Server

- Run: `go run cmd/fiubados-server/server.go <port_number>`
- Shutdown: type `EXIT` on prompt
- Closes when client sends `STOP` message
- Accepts multiple connections

## Client

- Run: `go run cmd/fiubados-client/client.go <ip:port_number>`
- Stop: type `STOP` on prompt
