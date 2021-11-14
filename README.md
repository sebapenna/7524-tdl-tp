# 7524 - Teoría del Lenguaje: Go

- Run: `go run .`
- Cleanup unused dependencies: `go mod tidy`

## Server

- Run: `go run server.go <port_number>`
- Closes when client sends `STOP` message

## Client

- Run: `go run server.go <ip:port_number>`
- Stop: type `STOP` on prompt
