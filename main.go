package main

import (
	"os"

	"github.com/sebapenna/7524-tdl-tp/client"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"github.com/sebapenna/7524-tdl-tp/server"
)

const (
	ExpectedArgs           = 3
	ConnectionInfoPosition = 1
	DesiredRolePosition    = 2
	clientRole             = "client"
	serverRole             = "server"
)

func main() {
	arguments := os.Args
	if len(arguments) != ExpectedArgs {
		logger.LogErrorMessage("Wrong number of arguments")
		return
	}

	if arguments[DesiredRolePosition] == clientRole {
		client.RunClient(arguments[ConnectionInfoPosition])
	} else if arguments[DesiredRolePosition] == serverRole {
		server.RunServer(arguments[ConnectionInfoPosition])
	}

}
