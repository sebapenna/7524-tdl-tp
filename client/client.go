package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
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

	reader := bufio.NewReader(os.Stdin)

	for {

		messageFromServer, err := common.Receive(currentSocket)
		if err != nil {
			fmt.Println("Server disconnected. Client exiting...")
			return
		}

		if messageFromServer == CloseConnectionCommand {
			fmt.Println("Client exiting...")
			return
		}

		fmt.Println("->: " + messageFromServer)

		fmt.Print(">> ")
		textFromPrompt, _ := reader.ReadString('\n')

		if textFromPrompt == CloseConnectionCommand {
			fmt.Println("Client exiting...")
			return
		}

		common.Send(currentSocket, textFromPrompt)

	}

}
