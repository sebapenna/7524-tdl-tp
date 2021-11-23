package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
)

// Runs server actions before starting the game
func HandshakeServer(player Player) bool {
	return startUpMenuServer(player)
}

// Runs client actions before starting the game
func HandshakeClient(currentSocket net.Conn) bool {

	for {
		reader := bufio.NewReader(os.Stdin)
		messageFromServer, err := common.Receive(currentSocket)
		if err != nil {
			logger.LogInfo("Server disconnected. Client exiting...")
			return false
		}

		if messageFromServer == CloseConnectionCommand {
			logger.LogInfo("Client exiting...")

			return false
		}

		if strings.HasPrefix(messageFromServer, common.WelcomeMessage) {
			logger.PrintMessageReceived(messageFromServer)
			common.Send(currentSocket, common.Success)
			messageFromServerAux, err := common.Receive(currentSocket)
			if err != nil {
				logger.LogInfo("Server disconnected. Client exiting...")
				return false
			}
			logger.PrintMessageReceived(messageFromServerAux)
			common.Send(currentSocket, common.Success)
			messageFromServerAux2, err := common.Receive(currentSocket)
			if err != nil {
				logger.LogInfo("Server disconnected. Client exiting...")
				return false
			}
			logger.PrintMessageReceived(messageFromServerAux2)
		} else if strings.HasPrefix(messageFromServer, common.HelpMessage) {
			logger.PrintMessageReceived(messageFromServer)
			common.Send(currentSocket, common.Success)
			messageFromServerAux, err := common.Receive(currentSocket)
			if err != nil {
				logger.LogInfo("Server disconnected. Client exiting...")
				return false
			}
			logger.PrintMessageReceived(messageFromServerAux)
		} else {
			logger.PrintMessageReceived(messageFromServer)
		}
		fmt.Print(">> ")
		textFromPrompt, _ := reader.ReadString('\n')
		common.Send(currentSocket, textFromPrompt)
	}
}

//returns true if after menu it is able to star looking for a game match , contrary case return false
func startUpMenuServer(player Player) bool {

	isAbleToLookForMatch := false
	for !isAbleToLookForMatch {
		messageFromClient, err := sendMainMenuOptions(player)
		if err != nil {
			logger.LogError(err)
			return false
		}
		switch messageFromClient {
		case common.OptionOne:
			logger.LogInfo("Player", player.id, "selected option 1, searching match...")
			isAbleToLookForMatch = true
			// ... //
		case common.OptionTwo:
			err = sendHelpSubMenuOptions(player)
			if err != nil {
				logger.LogError(err)
				return false
			}
		case common.OptionThree:
			disconnectPlayerFromMenu(player)
			return false
		}

	}
	return isAbleToLookForMatch

}

//Shows menu options and asks the client to pick one.
func sendMainMenuOptions(player Player) (string, error) {

	defer logger.LogInfo("Player", player.id, "redirected to main menu")

	// greets user and shows menu
	common.Send(player.socket, common.WelcomeMessage+strconv.Itoa(player.id))
	common.Receive(player.socket)
	common.Send(player.socket, common.ObjectiveMessage)
	common.Receive(player.socket)
	common.Send(player.socket, common.MainMenuOptions)
	// receives its answer
	messageFromClient, err := common.Receive(player.socket)
	return messageFromClient, err

}

// Shows options from HELP Submenu and asks the client to pick one.
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
		if messageFromClient == common.OptionOne {
			returnToMainMenu = true
		}
	}
	return err
}

// disconnect client from player that requested option 3 (Exit) from Menu.
func disconnectPlayerFromMenu(player Player) {
	common.Send(player.socket, CloseConnectionCommand)
	logger.LogInfo("Player", player.id, "disconnected")
}
