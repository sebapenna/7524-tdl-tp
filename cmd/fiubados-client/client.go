package main

import (
	"os"

	"github.com/sebapenna/7524-tdl-tp/client"
	"github.com/sebapenna/7524-tdl-tp/logger"
)

const (
	ClientExpectedArgs = 2
	ConnectionPosition = 1
)

func main() {
	arguments := os.Args
	if len(arguments) != ClientExpectedArgs {
		logger.LogErrorMessage("Wrong number of arguments")
		return
	}

	client.RunClient(arguments[ConnectionPosition])
}
