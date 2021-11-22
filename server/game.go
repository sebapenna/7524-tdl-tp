package server

import (
	"github.com/sebapenna/7524-tdl-tp/common"
	"github.com/sebapenna/7524-tdl-tp/logger"
	"math/rand"
	"strconv"
)

const (
	GameDuration = 10
)

// Game is responsible for handling a game between 2 players
type Game struct {
	player1 Player
	player2 Player
}

func RunStartGameAction(game Game) {
	logger.LogInfo("Starting a new game with player " + strconv.Itoa(game.player1.id) + " and player " + strconv.Itoa(game.player2.id))
	NotifyPlayersStartOfGame(game.player1, game.player2)
	questions := CreateRandomQuestions()

	for i := 0; i < GameDuration; i++ {
		randomQuestion := rand.Intn(len(questions))
		questionToAsk := questions[randomQuestion]
		questions[randomQuestion] = questions[len(questions)-1]
		questions = questions[:len(questions)-1]
		questionToAsk.questionNumber = i
		logger.LogInfo("Question to ask:", questionToAsk)
		logger.LogInfo(strconv.Itoa(len(questions)))
		AskQuestionToPlayers(game.player1, game.player2, questionToAsk)
	}
}

func NotifyPlayersStartOfGame(player1 Player, player2 Player) {
	common.Send(player1.socket, "Starting game! You are going to play against: "+strconv.Itoa(player2.id)+". Press enter to start")
	common.Send(player2.socket, "Starting game! You are going to play against: "+strconv.Itoa(player1.id)+". Press enter to start")
}

func AskQuestionToPlayers(player1 Player, player2 Player, question Question) {
	common.Send(player1.socket, "Question "+strconv.Itoa(question.questionNumber)+": "+question.question+" 1)"+question.options[0]+" 2)"+question.options[1]+" 3)"+question.options[2])
	common.Send(player2.socket, "Question "+strconv.Itoa(question.questionNumber)+": "+question.question+" 1)"+question.options[0]+" 2)"+question.options[1]+" 3)"+question.options[2])

	responsePlayer1, err := common.Receive(player1.socket)
	if err != nil {
		logger.LogError(err)
	}
	logger.LogInfo("Player 1 answer:", responsePlayer1)

	responsePlayer2, err := common.Receive(player2.socket)
	if err != nil {
		logger.LogError(err)
	}
	logger.LogInfo("Player 2 answer:", responsePlayer2)
}
