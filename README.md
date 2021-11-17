# 7524 - Teoría del Lenguaje: Go

- Cleanup unused dependencies: `go mod tidy`

## Server

- Run: `go run main.go <port_number> server`
- Shutdown: type `EXIT` on prompt
- Closes when client sends `STOP` message
- Accepts multiple connections

## Client

- Run: `go run main.go <ip:port_number> client`
- Stop: type `STOP` on prompt
