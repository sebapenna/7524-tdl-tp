package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"github.com/sebapenna/7524-tdl-tp/server"
)

const (
	ConnectionType = "tcp"
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

	continueGame := server.HandshakeClient(currentSocket)

	if continueGame == false {
		return
	}
	runClientGameLoop(currentSocket)
}

// Runs Client actions in game
func runClientGameLoop(currentSocket net.Conn) {

	promptReader := bufio.NewReader(os.Stdin)
	for {

		messageFromServer, err := common.Receive(currentSocket)
		if err != nil {
			logger.LogInfo(common.DisconnectAndExitMessage)
			return
		}

		if messageFromServer == common.CloseConnectionCommand {
			logger.LogInfo(common.ExitMessage)
			return
		}

		logger.PrintMessageReceived(messageFromServer)
		fmt.Print(">> ")
		textFromPrompt, _ := promptReader.ReadString('\n')

		if textFromPrompt == common.CloseConnectionCommand {
			logger.LogInfo(common.ExitMessage)
			return
		}

		common.Send(currentSocket, textFromPrompt)
	}

}
