package main

import (
	"os"

	"github.com/sebapenna/7524-tdl-tp/logger"
	"github.com/sebapenna/7524-tdl-tp/server"
)

const (
	ServerExpectedArgs = 2
	PortPos            = 1
)

func main() {
	arguments := os.Args
	if len(arguments) != ServerExpectedArgs {
		logger.LogErrorMessage("Wrong number of arguments")
		return
	}

	server.RunServer(arguments[PortPos])
}
