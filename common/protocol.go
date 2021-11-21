package common

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	WelcomeMessage  = "Welcome to FIUBADOS Player: "
	MainMenuOptions = "(1) Play  (2) Help  (3) Exit"
	// Le puse todos esos espacios al HELP para que se impriman las instrucciones una abajo de la otra
	HelpMessage            = "HELP:                                                                             ~~~ This game consists of 10 multiple choice questions ~~~                            ~~~ Each player will answer the number of the option chosen ~~~                       ~~~ If the first player to answer does it correctly gets the point ~~~                ~~~ If not the other player does ~~~                                                  ~~~ At the end of the game the player with the highest score wins ~~~"
	HelpMenuOptions        = "(1) Back to Main Menu"
	OptionOne              = "1"
	OptionTwo              = "2"
	OptionThree            = "3"
	CloseConnectionCommand = "STOP"
)

// Runs Client actions in game
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
