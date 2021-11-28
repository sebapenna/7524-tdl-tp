package server

import (
	"net"

	"github.com/sebapenna/7524-tdl-tp/logger"
)

// Lobby is in charge of handling incoming connection
// requests from clients and managing the players and games
type Lobby struct {
	listenSocket net.Listener
	players      []Player
	games        []Game
}

// RunLobby Sets the lobby to run
// Will keep receiving connections from clients until
// the server is shutdown.
func RunLobby(lobby Lobby) {
	//Channel where players will inform Lobby when they are ready to play
	channel := make(chan Player)
	go CreateGames(channel)
	for {
		/* Accept new connections or handle error if socket disconnected */
		currentSocket, err := lobby.listenSocket.Accept()
		if err != nil {
			logger.LogError(err)
			logger.LogInfo("Server shutdown")
			return
		}
		/* Create new player and save it into the already existing ones */
		newPlayer := Player{id: len(lobby.players) + 1, socket: currentSocket, channelPlayersReadyToPlay: channel}
		lobby.players = append(lobby.players, newPlayer)
		logger.LogInfo("Connection accepted: player", newPlayer.id)

		/* Create a new thread for the latest player */
		go RunPlayerAction(newPlayer)
	}
}

//Check players that are ready to play and starts a game
func CreateGames(channelPlayersReadyToPlay chan Player) {
	var playersReady []Player
	for {
		player := <-channelPlayersReadyToPlay
		playersReady = append(playersReady, player)
		logger.LogInfo("Player", player.id, "is ready to play!")
		if len(playersReady) == 2 {
			newGame := Game{player1: playersReady[0], player2: playersReady[1]}
			playersReady = append(playersReady[:0], playersReady[1:]...)
			playersReady = append(playersReady[:0], playersReady[1:]...)
			go RunStartGameAction(newGame)
		}
	}
}

// ShutdownLobby Shuts down the lobby by closing the listen socket
// and disconnecting every existing player
func ShutdownLobby(lobby Lobby) {
	/* Close socket receiving connections */
	lobby.listenSocket.Close()
	/* Disconnect every player */
	for i := range lobby.players {
		DisconnectPlayer(lobby.players[i])
	}
}
