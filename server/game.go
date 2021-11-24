package server

import (
	"errors"
	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"math/rand"
	"strconv"
	"time"
)

const (
	GameDuration = 10
)

// Game is responsible for handling a game between 2 players
type Game struct {
	player1 Player
	player2 Player
}

func DisconnectGame(game Game) {
	DisconnectPlayer(game.player1)
	DisconnectPlayer(game.player2)
}

func RunStartGameAction(game Game) {
	defer DisconnectGame(game)

	logger.LogInfo("Starting a new game with player " + strconv.Itoa(game.player1.id) + " and player " + strconv.Itoa(game.player2.id))
	err := notifyPlayersStartOfGame(game.player1, game.player2)
	if err != nil {
		logger.LogError(err)
		return
	}

	questions := CreateRandomQuestions()

	rand.Seed(time.Now().UnixNano()) // Set random seed to randomize questions
	for i := 0; i < GameDuration; i++ {
		randomQuestion := rand.Intn(len(questions))
		questionToAsk := questions[randomQuestion]

		questions[randomQuestion] = questions[len(questions)-1]
		questions = questions[:len(questions)-1]
		questionToAsk.questionNumber = i

		askQuestionToPlayers(game.player1, game.player2, questionToAsk)
	}
}

func readyToPlayLoop(player Player, otherPlayer Player, readyChannel chan bool) {
	msgToSend := func(id int) string {
		return "You are playing against player " + strconv.Itoa(id) + ". Enter READY when ready to play"
	}

	for {
		common.Send(player.socket, msgToSend(otherPlayer.id))
		msg, err := common.Receive(player.socket)
		if err != nil {
			logger.LogError(err)
			readyChannel <- false // Error in connection, return error
			return
		}
		if msg == common.ReadyToPlay {
			break
		}
	}

	readyChannel <- true
}

func notifyPlayersStartOfGame(player1 Player, player2 Player) error {
	/*
		Bool channel to notify player is ready to play.
		True means players is ready to play. False that there was an
		error and the game can not start.
	*/
	readyChannel := make(chan bool)

	go readyToPlayLoop(player1, player2, readyChannel)
	go readyToPlayLoop(player2, player1, readyChannel)

	playersReady := 0
	for {
		ready := <-readyChannel // New player notified that is ready
		if !ready {
			// todo: mejora, notificar al player que no se desconecto que el otro abandono la partida
			return errors.New("player disconnected before game started")
		}
		playersReady++
		if playersReady == 2 {
			break
		}
	}
	return nil
}

func askQuestionToPlayers(player1 Player, player2 Player, question Question) {
	msgToSend := func(question Question) string {
		return "Question " + strconv.Itoa(question.questionNumber) + ": " + question.question + " 1)" + question.options[0] + " 2)" + question.options[1] + " 3)" + question.options[2]
	}

	common.Send(player1.socket, msgToSend(question))
	common.Send(player2.socket, msgToSend(question))

	responsePlayer1, err := common.Receive(player1.socket)
	// todo: si hay un error hay que desonectar a los 2
	if err != nil {
		logger.LogError(err)
	}
	logger.LogInfo("Player 1 answer:", responsePlayer1)

	responsePlayer2, err := common.Receive(player2.socket)
	// todo: si hay un error hay que desonectar a los 2
	if err != nil {
		logger.LogError(err)
	}
	logger.LogInfo("Player 2 answer:", responsePlayer2)
}
