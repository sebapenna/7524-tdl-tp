package server

import (
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	CloseConnectionCommand = "STOP"
)

// Player represents each player connected to the server
type Player struct {
	id     int
	socket net.Conn
}

// DisconnectPlayer Closes the connection of the current's
// player client
func DisconnectPlayer(player Player) {
	fmt.Println("Disconnecting player " + strconv.Itoa(player.id))
	player.socket.Close()
}

// RunPlayerAction starts listening incoming requests
// from the client linked to the player and managing
// the game
func RunPlayerAction(player Player) {
	/* Disconnect player when loop finished */
	defer DisconnectPlayer(player)

	for {
		netData, err := common.Receive(player.socket)
		if err != nil {
			logger.LogError(err)
			return
		}
		if strings.TrimSpace(string(netData)) == CloseConnectionCommand {
			fmt.Println("Client disconnected")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		common.Send(player.socket, t.Format(time.RFC3339))
	}
}
