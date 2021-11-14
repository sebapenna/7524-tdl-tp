package client

import (
	"bufio"
	"fmt"
	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"net"
	"os"
	"strings"
)

const (
	ConnectionType         = "tcp"
	CloseConnectionCommand = "STOP"
)

func RunClient(connection string) {
	c, err := net.Dial(ConnectionType, connection)
	if err != nil {
		logger.LogError(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		common.Send(c, text)

		message, _ := common.Receive(c)
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == CloseConnectionCommand {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}
