package server

import (
	"bufio"
	"net"
	"os"
	"strings"

	"github.com/sebapenna/7524-tdl-tp/logger"
)

const (
	ConnectionType        = "tcp"
	ShutdownServerCommand = "EXIT"
)

func formatPort(port string) string {
	return ":" + port
}

// Shuts down the server if the specified command is read
// by the server input
func shutdownServer(lobby Lobby) {
	var requestToShutdownWasMade bool
	for !requestToShutdownWasMade {
		/* Keep reading stdin until the shutdown command is found */
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(input) == ShutdownServerCommand {
			/* Shutdown the lobby and break the loop */
			requestToShutdownWasMade = true
			ShutdownLobby(lobby)

		}
	}
}

// RunServer starts the server and enable incoming
// connections to be handled
func RunServer(port string) {
	lobbySocket, err := net.Listen(ConnectionType, formatPort(port))
	logger.LogInfo("Server listening on port " + port)
	if err != nil {
		logger.LogError(err)
		return
	}
	lobby := Lobby{listenSocket: lobbySocket, players: []Player{}, games: []Game{}}

	/* Create thread to shut down server whenever it's requested */
	go shutdownServer(lobby)

	/* Put the lobby to work in the current thread */
	RunLobby(lobby)
}
