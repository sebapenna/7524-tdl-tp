package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
)

// Runs server actions before starting the game
// returns true if the handshake with the client was successful; false otherwise
func HandshakeServer(player Player) bool {

	return startUpMenuServer(player)

}

// Runs client actions before starting the game
// returns true if the handshake with the server was successful; false otherwise
func HandshakeClient(currentSocket net.Conn) bool {

	return startUpMenuClient(currentSocket)

}

//(SERVER-SIDE) Sends messages to the client and analizes the answers received from it to display the corresponding Menu section / options.
//returns true if after menu it is able to star looking for a game match , otherwise returns false
func startUpMenuServer(player Player) bool {

	logger.LogInfo("Player", player.id, "directed to main menu")

	isAbleToLookForMatch := false
	for !isAbleToLookForMatch {
		messageFromClient, err := sendMainMenuOptions(player)
		if err != nil {
			logger.LogError(err)
			return false
		}

		switch messageFromClient {
		case common.PlayOption:
			logger.LogInfo("Player", player.id, "selected option 1, searching match...")
			isAbleToLookForMatch = true
			sendFindingMatchMessage(player)

		case common.HelpOption:
			err = sendHelpSubMenuOptions(player)
			if err != nil {
				logger.LogError(err)
				return false
			}

		case common.ExitOption:
			disconnectPlayerFromMenu(player)
			return false
		}

	}
	return isAbleToLookForMatch

}

//(CLIENT-SIDE) Sends messages to the server and analizes the answers received from it to select the desired Menu section / options.
//returns true if after menu it is able to star looking for a game match , otherwise returns false
func startUpMenuClient(currentSocket net.Conn) bool {

	for {

		messageFromServer, err := common.Receive(currentSocket)
		VerifyErrorReveivedFromServer(err)

		if messageFromServer == CloseConnectionCommand {
			logger.LogInfo(common.ExitMessage)
			return false
		}

		if strings.HasPrefix(messageFromServer, common.WelcomeMessage) {
			printWelcomeReceivedFromServer(currentSocket)

		} else if strings.HasPrefix(messageFromServer, common.HelpMessage) {
			printHelpReceivedFromServer(currentSocket)

		} else {
			logger.PrintMessageReceived(messageFromServer)
		}

		if messageFromServer == common.SearchingMatchMessage {
			common.Send(currentSocket, common.Success)

		} else {
			readFromPromptAndSendItToTheServer(currentSocket)
		}

	}

}

//(SERVER-SIDE) Sends menu options to the client and asks it to pick one.
func sendMainMenuOptions(player Player) (string, error) {

	// greets user and shows menu
	common.Send(player.socket, common.WelcomeMessage+player.name)
	common.Receive(player.socket)
	common.Send(player.socket, common.ObjectiveMessage)
	common.Receive(player.socket)
	common.Send(player.socket, common.MainMenuOptions)
	// receives its answer
	messageFromClient, err := common.Receive(player.socket)
	return messageFromClient, err

}

//(SERVER-SIDE) Sends HELP submenu options to the client and asks it to pick one.
func sendHelpSubMenuOptions(player Player) error {

	logger.LogInfo("Player", player.id, "selected option 2, showing help...")
	defer logger.LogInfo("Player", player.id, "redirected to main menu")

	var (
		messageFromClient string
		err               error
		returnToMainMenu  bool
	)

	for !returnToMainMenu {
		common.Send(player.socket, common.HelpMessage)
		common.Receive(player.socket)
		common.Send(player.socket, common.HelpMenuOptions)
		messageFromClient, err = common.Receive(player.socket)
		if messageFromClient == common.PlayOption {
			returnToMainMenu = true
		}
	}
	return err
}

//(SERVER-SIDE) Sends a message to the client telling it that it's currently in queue looking for a match.
func sendFindingMatchMessage(player Player) {
	common.Send(player.socket, common.SearchingMatchMessage)
	common.Receive(player.socket)
}

//(SERVER-SIDE) Sends a message for the client to disconnect when it requested option 3 (Exit) from the main Menu.
func disconnectPlayerFromMenu(player Player) {
	common.Send(player.socket, CloseConnectionCommand)
	logger.LogInfo("Player", player.id, "disconnected")
}

//(CLIENT-SIDE) Executes panic if there was a problem receiving a message from the server.
//Does nothing if there weren't any problems.
func VerifyErrorReveivedFromServer(err error) {
	if err != nil {
		panic(common.DisconnectAndExitMessage)
	}
}

//(CLIENT-SIDE) Finishes the send-receive protocol in order to print the entire main Menu with its options
func printWelcomeReceivedFromServer(currentSocket net.Conn) {
	logger.PrintMessageReceived(common.WelcomeMessage)
	common.Send(currentSocket, common.Success)
	messageFromServerAux, err := common.Receive(currentSocket)
	VerifyErrorReveivedFromServer(err)

	logger.PrintMessageReceived(messageFromServerAux)
	common.Send(currentSocket, common.Success)
	messageFromServerAux2, err := common.Receive(currentSocket)
	VerifyErrorReveivedFromServer(err)

	logger.PrintMessageReceived(messageFromServerAux2)
}

//(CLIENT-SIDE) Finishes the send-receive protocol in order to print the entire HELP submenu with its options
func printHelpReceivedFromServer(currentSocket net.Conn) {
	logger.PrintMessageReceived(common.HelpMessage)
	common.Send(currentSocket, common.Success)
	messageFromServerAux, err := common.Receive(currentSocket)
	VerifyErrorReveivedFromServer(err)

	logger.PrintMessageReceived(messageFromServerAux)
}

//(CLIENT-SIDE) gets the next message to send to the server from the client's prompt
func readFromPromptAndSendItToTheServer(currentSocket net.Conn) {
	promptReader := bufio.NewReader(os.Stdin)
	fmt.Print(string(common.ColorGreen), ">> ", string(common.ColorReset))
	textFromPrompt, _ := promptReader.ReadString('\n')
	common.Send(currentSocket, textFromPrompt)
}
