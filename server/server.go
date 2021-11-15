package server

import (
	"bufio"
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	"os"
	"strings"
)

const (
	ConnectionType        = "tcp"
	ShutdownServerCommand = "EXIT"
)

func formatPort(port string) string {
	return ":" + port
}

func shutdownServer(lobby Lobby) {
	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(input) == ShutdownServerCommand {
			ShutdownLobby(lobby)
			break
		}
	}
}

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

	RunLobby(lobby)

}
