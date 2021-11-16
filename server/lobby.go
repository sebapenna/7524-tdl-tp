package server

import (
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	"strconv"
)

// Lobby is in charge of handling incoming connection
// requests from clients and managing the players and games
type Lobby struct {
	listenSocket net.Listener
	players      []Player
}

// RunLobby Sets the lobby to run
// Will keep receiving connections from clients until
// the server is shutdown.
func RunLobby(lobby Lobby) {
	var nextPlayerId = 0
	for {
		/* Accept new connections or handle error if socket disconnected */
		c, err := lobby.listenSocket.Accept()
		if err != nil {
			logger.LogError(err)
			fmt.Println("Server shutdown")
			return
		}

		/* Create new player and save it into the already existing ones */
		newPlayer := Player{nextPlayerId, c}
		lobby.players = append(lobby.players, newPlayer)
		fmt.Println("Connection accepted: player " + strconv.Itoa(nextPlayerId))
		nextPlayerId += 1

		/* Create a new thread for the latest player */
		go RunPlayerAction(newPlayer)
	}
}

// ShutdownLobby Shuts down the lobby by closing the listen socket
// and disconnecting every existing player
func ShutdownLobby(lobby Lobby) {
	/* Close socket receiving connections */
	lobby.listenSocket.Close()

	/* Disconnect every player */
	for i := range lobby.players {
		DisconnectPlayer(lobby.players[i])
	}
}
