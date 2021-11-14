package server

import (
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	"strings"
	"time"
)

const (
	ConnectionType         = "tcp"
	CloseConnectionCommand = "STOP"
)

func formatPort(port string) string {
	return ":" + port
}

func handleConnection(clientSocket net.Conn) {
	for {
		netData, err := common.Receive(clientSocket)
		if err != nil {
			logger.LogError(err)
			return
		}
		if strings.TrimSpace(string(netData)) == CloseConnectionCommand {
			fmt.Println("Exiting TCP server!")
			clientSocket.Close()
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		common.Send(clientSocket, t.Format(time.RFC3339))
	}
}

func RunServer(port string) {
	l, err := net.Listen(ConnectionType, formatPort(port))
	fmt.Println("Server listening on port " + port)
	if err != nil {
		logger.LogError(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		fmt.Println("Connection accepted")
		if err != nil {
			logger.LogError(err)
			return
		}
		go handleConnection(c)
	}

}
