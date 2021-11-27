package server

import (
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	//"time"
)

const (
	CloseConnectionCommand = "STOP"
)

// Player represents each player connected to the server
type Player struct {
	id                        int
	name                      string
	socket                    net.Conn
	points                    int
	channelPlayersReadyToPlay chan<- Player
}

// DisconnectPlayer Closes the connection of the current's
// player client
func DisconnectPlayer(player Player) {
	logger.LogInfo("Disconnecting player", player.id)
	player.socket.Close()
}

// RunPlayerAction starts listening incoming requests
// from the client linked to the player and managing
// the game
func RunPlayerAction(player Player) {
	readyToSearchForGame := HandshakeServer(player) /*StartUpMenu(player)*/

	if readyToSearchForGame {
		player.channelPlayersReadyToPlay <- player
	}
}
