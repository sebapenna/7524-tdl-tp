package server

import (
	"fmt"
	"net"
	//"strconv"

	//"github.com/sebapenna/7524-tdl-tp/logger"
)

// Game is responsible for handling a game between 2 players
type Game struct {
	listenSocket net.Listener
	player1      Player
	player2      Player
}

func RunStartGameAction(game Game) {
    fmt.Println("Starting a new game with player ", game.player1.id, " and player ", game.player2.id)

}