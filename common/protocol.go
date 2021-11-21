package common

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	WelcomeMessage         = "Bienvenido a FIUBADOS Jugador:"
	MainMenuOptions        = "(1) Play  (2) Help  (3) Exit"
	HelpMessage            = "AYUDA: *texto de ayuda * "
	HelpMenuOptions        = "(1) Back to Main Menu "
	OptionOne              = "1"
	OptionTwo              = "2"
	OptionThree            = "3"
	CloseConnectionCommand = "STOP"
)

func RunClientProtocol(currentSocket net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {

		messageFromServer, err := Receive(currentSocket)
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

		Send(currentSocket, textFromPrompt)

	}

}
