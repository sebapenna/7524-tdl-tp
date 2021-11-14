package server

import (
	"bufio"
	"fmt"
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

func RunServer(port string) {
	l, err := net.Listen(ConnectionType, formatPort(port))
	fmt.Println("Server listening on port " + port)
	if err != nil {
		logger.LogError(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	fmt.Println("Connection accepted")
	if err != nil {
		logger.LogError(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			logger.LogError(err)
			return
		}
		if strings.TrimSpace(string(netData)) == CloseConnectionCommand {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}

}
