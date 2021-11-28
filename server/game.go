package server

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
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

	runGameLoop(game.player1, game.player2)
}

func readyToPlayLoop(player Player, otherPlayer Player, readyChannel chan bool) {
	msgToSend := func(playerName string) string {
		return "You are playing against player " + playerName + ". Enter READY when ready to play"
	}

	var playerIsReady bool
	for playerIsReady == false {
		common.Send(player.socket, msgToSend(otherPlayer.name))
		msg, err := common.Receive(player.socket)
		if err != nil {
			logger.LogError(err)
			readyChannel <- false // Error in connection, return error
			return
		}
		if msg == common.ReadyToPlay {
			playerIsReady = true
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

type Answer struct {
	player       *Player
	question     Question
	optionChosen int
}

func sendQuestionsAndReceiveAnswers(
	player *Player,
	questionsChannel chan Question,
	answerChannel chan Answer,
) {
	msgToSend := func(question Question) string {
		return "Question " + strconv.Itoa(question.questionNumber) + ": " + question.questionStr + " 1)" + question.options[0] + " 2)" + question.options[1] + " 3)" + question.options[2]
	}

	for question := range questionsChannel {
		common.Send(player.socket, msgToSend(question))
		optionChosen, err := common.Receive(player.socket)

		// todo: si hay un error hay que desconectar a los 2
		if err != nil {
			logger.LogError(err)
		}

		opt, _ := strconv.Atoi(optionChosen)
		answerChannel <- Answer{
			player:       player,
			question:     question,
			optionChosen: opt,
		}
	}
}

func notifyWinner(player1 Player, player2 Player) {
	notifyGameResult := func(msg string, player1 Player, player2 Player) {
		common.Send(player1.socket, msg)
		common.Send(player2.socket, msg)
	}

	switch {
	case player1.points > player2.points:
		notifyGameResult("Player "+player1.name+" won! Thanks for playing FIUBADOS! Press any key to exit", player1, player2)
	case player2.points > player1.points:
		notifyGameResult("Player "+player2.name+" won! Thanks for playing FIUBADOS! Press any key to exit", player1, player2)
	default:
		notifyGameResult("Game tied! Thanks for playing FIUBADOS! Press any key to exit", player1, player2)
	}
}

func runGameLoop(player1 Player, player2 Player) {
	questionsChannel1 := make(chan Question)
	questionsChannel2 := make(chan Question)
	answersChannel := make(chan Answer)
	defer close(questionsChannel1)
	defer close(questionsChannel2)
	defer close(answersChannel)

	go sendQuestionsAndReceiveAnswers(&player1, questionsChannel1, answersChannel)
	go sendQuestionsAndReceiveAnswers(&player2, questionsChannel2, answersChannel)

	questions := CreateRandomQuestions()

	rand.Seed(time.Now().UnixNano()) // Set random seed to randomize questions
	for i := 1; i <= GameDuration; i++ {
		randomQuestion := rand.Intn(len(questions))
		questionToAsk := questions[randomQuestion]

		questions[randomQuestion] = questions[len(questions)-1]
		questions = questions[:len(questions)-1]
		questionToAsk.questionNumber = i
		questionsChannel1 <- questionToAsk
		questionsChannel2 <- questionToAsk

		answer1 := <-answersChannel
		answer2 := <-answersChannel

		if answer1.optionChosen == questionToAsk.correctOption {
			answer1.player.points++
		} else {
			answer2.player.points++
		}
	}

	notifyWinner(player1, player2)
}
