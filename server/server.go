package server

import (
	"bufio"
	"fmt"
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
	for {
		/* Keep reading stdin until the shutdown command is found */
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(input) == ShutdownServerCommand {
			/* Shutdown the lobby and break the loop */
			ShutdownLobby(lobby)
			break
		}
	}
}

// RunServer starts the server and enable incoming
// connections to be handled
func RunServer(port string) {
	l, err := net.Listen(ConnectionType, formatPort(port))
	fmt.Println("Server listening on port " + port)
	if err != nil {
		logger.LogError(err)
		return
	}
	lobby := Lobby{l, []Player{}}

	/* Create thread to shut down server */
	go shutdownServer(lobby)

	/* Put the lobby to work in the current thread */
	RunLobby(lobby)
}
