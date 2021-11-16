package server

import (
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
)

type Lobby struct {
	listenSocket net.Listener
	players      []Player
}

func RunLobby(lobby Lobby) {
	var nextPlayerId = 0
	for {
		c, err := lobby.listenSocket.Accept()
		if err != nil {
			logger.LogError(err)
			fmt.Println("Server shutdown")
			return
		}

		newPlayer := Player{nextPlayerId, c}
		lobby.players = append(lobby.players, newPlayer)
		nextPlayerId += 1
		fmt.Println("Connection accepted")

		go RunPlayerAction(newPlayer)
	}
}

func ShutdownLobby(lobby Lobby) {
	/* Close socket receiving connections */
	lobby.listenSocket.Close()

	/* Disconnect every player */
	for i := range lobby.players {
		DisconnectPlayer(lobby.players[i])
	}
}
