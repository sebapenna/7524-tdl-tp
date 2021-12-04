# 7524 - Teoría del Lenguaje: Go.
# FIUBADOS.

## Bienvenidos a _FIUBADOS_.
El objetivo de este juego es introducir a los alumnos que ingresan a la facultad en diversas cuestiones administrativas / datos curiosos, siendo en el futuro útil para el desarrollo de su carrera profesional.

Este juego consiste en partidas 1vs1 en las que dos jugadores responden varias preguntas de opción múltiple.
Cada jugador contestará el número de la opción que considere correcta en cada pregunta.

El jugador que responda de forma correcta aumenta su puntuación.
Si ambos jugadores responden correctamente una pregunta, el primero que haya respondido se lleva puntos adicionales. 
Asímismo, si solo uno de los jugadores responde correctamente, se lleva puntos adicionales.
El jugador que responda incorrectamente no suma ningun punto.

Al final del juego, el jugador con la mayor puntuación gana. ¡Buena suerte!

(Pueden llevarse a cabo multiples partidas a la vez).

[//]: # (- Cleanup unused dependencies: `go mod tidy`)

## Guía de instalación y ejecución.
### Instalación:
- En el directorio **_root_** del proyecto, correr los siguientes comandos:
```bash
go build ./cmd/... && go install ./cmd/...
```
### Ejecución:
#### Server:
- Con la App instalada:
```bash 
fiubados-server <port_number>
```
- Con la App SIN instalar:
```bash
go run cmd/fiubados-server/server.go <port_number>
```
#### Cliente:
- Con la App instalada:
```bash 
fiubados-client <address:port_number>
```
- Con la App SIN instalar:
```bash
go run cmd/fiubados-client/client.go <address:port_number>
```

## Server.
- Acepta conexiones de multiples jugadores a la vez.
- Puede ser desconectado enviandole `EXIT` desde su prompt. 

