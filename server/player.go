package server

import (
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	"strings"
	"time"
)

type Player struct {
	id     int
	socket net.Conn
}

func DisconnectPlayer(player Player) {
	player.socket.Close()
}

func RunPlayerAction(player Player) {
	for {
		netData, err := common.Receive(player.socket)
		if err != nil {
			logger.LogError(err)
			return
		}
		if strings.TrimSpace(string(netData)) == CloseConnectionCommand {
			fmt.Println("Client disconnected. Closing connection...")
			player.socket.Close()
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		common.Send(player.socket, t.Format(time.RFC3339))
	}
}
