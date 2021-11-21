package client

import (
	"net"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"github.com/sebapenna/7524-tdl-tp/server"
)

const (
	ConnectionType         = "tcp"
	CloseConnectionCommand = "STOP"
)

// RunClient connects to the server and keeps the connection
// alive while the game is active or the server is not
// closed
func RunClient(connection string) {
	currentSocket, err := net.Dial(ConnectionType, connection)
	if err != nil {
		logger.LogError(err)
		return
	}
	defer currentSocket.Close()

	continueGame := server.HandShakeClient(currentSocket)

	if continueGame == false {
		return
	}
	common.RunClientProtocol(currentSocket)
}
