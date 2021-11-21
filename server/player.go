package server

import (
	"fmt"
	"net"
	"strconv"
	//"time"
)

const (
	CloseConnectionCommand = "STOP"
)

// Player represents each player connected to the server
type Player struct {
	id int
	//name   string
	socket                   net.Conn
	puntaje                  int
	chanelPlayersReadyToPlay chan<- Player
	chanelQuestions          chan<- Question
}

// DisconnectPlayer Closes the connection of the current's
// player client
func DisconnectPlayer(player Player) {
	fmt.Println("Disconnecting player " + strconv.Itoa(player.id) /*+ " (" + player.name + ")"*/)
	player.socket.Close()
}

// RunPlayerAction starts listening incoming requests
// from the client linked to the player and managing
// the game
func RunPlayerAction(player Player) {

	puedeBuscarPartida := HandShakeServer(player) /*StartUpMenu(player)*/

	if puedeBuscarPartida {
		player.chanelPlayersReadyToPlay <- player

	}

}
