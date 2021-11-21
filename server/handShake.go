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
func HandShakeServer(player Player) bool {
	return startUpMenuServer(player)
}

// Runs client actions before starting the game
func HandShakeClient(currentSocket net.Conn) bool {

	for {
		reader := bufio.NewReader(os.Stdin)
		messageFromServer, err := common.Receive(currentSocket)
		if err != nil {
			fmt.Println("Server disconnected. Client exiting...")
			return false
		}
		if messageFromServer == CloseConnectionCommand {
			fmt.Println("Client exiting...")

			return false
		}
		if strings.HasPrefix(messageFromServer, common.WelcomeMessage) {
			fmt.Println("->: " + messageFromServer)
			common.Send(currentSocket, "")
			messageFromServerAux, err := common.Receive(currentSocket)
			if err != nil {
				fmt.Println("Server disconnected. Client exiting...")
				return false
			}
			fmt.Println("->: " + messageFromServerAux)
		} else if strings.HasPrefix(messageFromServer, common.HelpMessage) {
			fmt.Println("->: " + messageFromServer)
			common.Send(currentSocket, "")
			messageFromServerAux, err := common.Receive(currentSocket)
			if err != nil {
				fmt.Println("Server disconnected. Client exiting...")
				return false
			}
			fmt.Println("->: " + messageFromServerAux)
		} else {
			fmt.Println("->: " + messageFromServer)
		}
		fmt.Print(">> ")
		textFromPrompt, _ := reader.ReadString('\n')
		common.Send(currentSocket, textFromPrompt)
	}
}

//returns true if after menu it is able to star looking for a game match , contrary case return false
func startUpMenuServer(player Player) bool {

	isAbleToLookForMAtch := false
	for !isAbleToLookForMAtch {
		messageFromClient, err := sendMainMenuOptions(player)
		if err != nil {
			logger.LogError(err)
			return false
		}
		fmt.Println("-> ", messageFromClient)
		if messageFromClient == common.OptionOne {
			fmt.Println("Player " /*player.name*/, player.id, " selected option 1, searching match...")
			isAbleToLookForMAtch = true
			// ... //
		} else if messageFromClient == common.OptionTwo {
			err = sendHelpSubMenuOptions(player)
			if err != nil {
				logger.LogError(err)
				return false
			}
		} else if messageFromClient == common.OptionThree {
			disconnectPlayerFromMenu(player)
			return false
		}

	}
	return isAbleToLookForMAtch

}

//Shows menu options and asks the client to pick one.
func sendMainMenuOptions(player Player) (string, error) {

	defer fmt.Println("Player " /*player.name*/, player.id, " redirected to main menu")

	// greets user and shows menu
	common.Send(player.socket, common.WelcomeMessage+strconv.Itoa(player.id))
	common.Receive(player.socket)
	common.Send(player.socket, common.MainMenuOptions)
	// receives its answer
	messageFromClient, err := common.Receive(player.socket)
	return messageFromClient, err

}

// Shows options from HELP Submenu and asks the client to pick one.
func sendHelpSubMenuOptions(player Player) error {

	fmt.Println("Player" /*player.name*/, player.id, "selected option 2, showing help...")
	defer fmt.Println("Player" /*player.name*/, player.id, "redirected to main menu")

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
	fmt.Println("Player selected option 3, disconnecting client...")
	common.Send(player.socket, CloseConnectionCommand)
	fmt.Println("Client disconnected")
}
