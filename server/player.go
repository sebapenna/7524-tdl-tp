package server

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
)

const (
	CloseConnectionCommand = "STOP"
)

// Player represents each player connected to the server
type Player struct {
	id     int
	name   string
	socket net.Conn
}

// DisconnectPlayer Closes the connection of the current's
// player client
func DisconnectPlayer(player Player) {
	fmt.Println("Disconnecting player " + strconv.Itoa(player.id) + " (" + player.name + ")")
	player.socket.Close()
}

// RunPlayerAction starts listening incoming requests
// from the client linked to the player and managing
// the game
func RunPlayerAction(player Player) {
	/* Disconnect player when loop finished */
	defer DisconnectPlayer(player)

	puedeBuscarPartida := startupMenu(player)

	//que aca esté la logica de busqueda de partida en vez del for:
	for puedeBuscarPartida {
		t := time.Now()
		common.Send(player.socket, t.Format(time.RFC3339))

		messageFromClient, err := common.Receive(player.socket)
		if err != nil {
			logger.LogError(err)
			return
		}
		if strings.TrimSpace(messageFromClient) == CloseConnectionCommand {
			fmt.Println("Client disconnected")
			return
		}

		fmt.Print("-> ", messageFromClient)

	}

}

// Devuelve true si puede comenzar a buscar partida correctamente tras el menú , false en caso contrario
func startupMenu(player Player) bool {

	puedeBuscarPartida := false
	for !puedeBuscarPartida {

		messageFromClient, err := sendMainMenuOptions(player)
		if err != nil {
			logger.LogError(err)
			return false
		}
		fmt.Print("-> ", messageFromClient)

		if strings.TrimSpace(messageFromClient) == "1" {

			fmt.Print("Player ", player.name, " selected option 1, searching match...")
			puedeBuscarPartida = true
			// ... //

		} else if strings.TrimSpace(messageFromClient) == "2" {

			err = sendHelpSubMenuOptions(player)
			if err != nil {
				logger.LogError(err)
				return false
			}

		} else if strings.TrimSpace(messageFromClient) == "3" {

			disconnectPlayerFromMenu(player)
			return false

		}

	}

	return puedeBuscarPartida

}

//Muestra opciones del menú y le pide al cliente que elija una
func sendMainMenuOptions(player Player) (string, error) {

	defer fmt.Print("Player ", player.name, " redirected to main menu")

	// Saluda al usuario y le muestra el menú
	welcomeText := "Bienvenido a FIUBADOS:  (1) Play  (2) Help  (3) Exit"
	common.Send(player.socket, welcomeText)
	// Recibe su respuesta
	messageFromClient, err := common.Receive(player.socket)
	return messageFromClient, err

}

//Muestra opciones del Submenú de HELP y le pide al cliente que elija una
func sendHelpSubMenuOptions(player Player) error {

	fmt.Println("Player", player.name, "selected option 2, showing help...")

	defer fmt.Println("Player", player.name, "redirected to main menu")

	var (
		messageFromClient string
		err               error
		i                 int
		helpText          string
		volverAMainMenu   bool
	)

	textoDeAyudaPorPaginas := [3]string{"*texto de ayuda 1*", "*texto de ayuda 2*", "*texto de ayuda 3*"}

	for !volverAMainMenu {

		helpText = "AYUDA: " + textoDeAyudaPorPaginas[i] + " (1) Back to Main Menu (2) Next "
		common.Send(player.socket, helpText)

		messageFromClient, err = common.Receive(player.socket)

		if strings.TrimSpace(messageFromClient) == "1" {
			volverAMainMenu = true
		}

		if i < 2 {
			i++
		}

	}

	return err

}

// Desconecta al cliente del jugador que solicitó opcion 3 (Exit) del Menú
func disconnectPlayerFromMenu(player Player) {
	fmt.Println("Player selected option 3, disconnecting client...")
	common.Send(player.socket, CloseConnectionCommand)
	fmt.Println("Client disconnected")
}
