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
		return common.MatchingPlayerMessage + playerName + common.ReadyToPlayMessage
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

type PlayerError struct {
	player *Player
	err    error
}

func sendQuestionsAndReceiveAnswers(
	player *Player,
	questionsChannel chan Question,
	answerChannel chan Answer,
	errorChannel chan PlayerError,
) {
	msgToSend := func(question Question) string {
		return common.QuestionMessage + strconv.Itoa(question.questionNumber) + common.ColonMessage + question.questionStr + common.FirstOption + question.options[0] + common.SecondOption + question.options[1] + common.ThirdOption + question.options[2]
	}

	for question := range questionsChannel {
		var opt = 0
		for opt != 1 && opt != 2 && opt != 3 {
			errSend := common.Send(player.socket, msgToSend(question))
			if errSend != nil {
				logger.LogError(errSend)
				errorChannel <- PlayerError{player, errSend}
				return
			}

			optionChosen, errReceive := common.Receive(player.socket)

			if errReceive != nil {
				logger.LogError(errReceive)
				errorChannel <- PlayerError{player, errReceive}
				return
			}

			opt, _ = strconv.Atoi(optionChosen)
		}

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
		notifyGameResult(common.PlayerMessage+player1.name+common.WinnerMessage, player1, player2)
	case player2.points > player1.points:
		notifyGameResult(common.PlayerMessage+player2.name+common.WinnerMessage, player1, player2)
	default:
		notifyGameResult(common.TieMessage, player1, player2)
	}
}

func notifyOtherPlayerDisconnected(player Player) {
	common.Send(player.socket, common.OtherPlayerDisconnectedMessage)
}

// Handles player error
// Notifies the player that did not disconnect
// that it was the winner and then also disconnects it
func handlePlayerDisconnected(
	playerError PlayerError,
	player1 Player,
	player2 Player,
) {
	if playerError.player.id == player1.id {
		notifyOtherPlayerDisconnected(player2)
		DisconnectPlayer(player2)
	} else {
		notifyOtherPlayerDisconnected(player1)
		DisconnectPlayer(player1)
	}
}

// Reads the answers of every player
// and distributes the points based
// on them.
// Also checks if any error raises while
// players should have answered
func readAnswersAndDistributePoints(
	answersChannel chan Answer,
	errorChannel chan PlayerError,
	player1 *Player,
	player2 *Player,
	questionAsked Question,
) {
	// Create a loop to wait for both answers
	// but also keep checking if any error
	// raised in any player connection
	var answer1, answer2 Answer
	for questionAnswered := 0; questionAnswered < 2; questionAnswered++ {
		select {
		case playerError := <-errorChannel:
			handlePlayerDisconnected(playerError, *player1, *player2)
			return
		case answer := <-answersChannel:
			if questionAnswered == 0 {
				answer1 = answer // Save first answer received
			} else {
				answer2 = answer // Save second answer received
			}
		}
	}

	if answer1.optionChosen == questionAsked.correctOption {
		answer1.player.points++
	} else {
		answer2.player.points++
	}
}

func runGameLoop(player1 Player, player2 Player) {
	questionsChannel1 := make(chan Question)
	questionsChannel2 := make(chan Question)
	answersChannel := make(chan Answer)
	errorChannel := make(chan PlayerError)
	defer close(questionsChannel1)
	defer close(questionsChannel2)
	defer close(answersChannel)

	go sendQuestionsAndReceiveAnswers(&player1, questionsChannel1, answersChannel, errorChannel)
	go sendQuestionsAndReceiveAnswers(&player2, questionsChannel2, answersChannel, errorChannel)

	questions := CreateRandomQuestions()

	rand.Seed(time.Now().UnixNano()) // Set random seed to randomize questions
	for i := 1; i <= GameDuration; i++ {
		select {
		case playerError := <-errorChannel:
			handlePlayerDisconnected(playerError, player1, player2)
			return
		default:
			randomQuestion := rand.Intn(len(questions))
			questionToAsk := questions[randomQuestion]

			questions[randomQuestion] = questions[len(questions)-1]
			questions = questions[:len(questions)-1]
			questionToAsk.questionNumber = i
			questionsChannel1 <- questionToAsk
			questionsChannel2 <- questionToAsk

			readAnswersAndDistributePoints(answersChannel, errorChannel, &player1, &player2, questionToAsk)
		}
	}

	notifyWinner(player1, player2)
}
