package server

import (
	"net"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
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
	lastAnswerWasCorrect      bool
	wasFirstToAnswerCorrectly bool
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
	RequestPlayerName(&player)
	readyToSearchForGame := HandshakeServer(player)

	if readyToSearchForGame {
		player.channelPlayersReadyToPlay <- player
	}
}

//asks each player their name
func RequestPlayerName(player *Player) {
	common.Send(player.socket, common.AskForNameMessage)
	playerName, err := common.Receive(player.socket)

	player.name = playerName
	if err != nil {
		logger.LogError(err)
	}
}
