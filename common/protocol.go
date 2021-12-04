package common

const (
	AskForNameMessage = "Por favor, introduce un nombre para jugar: "
	WelcomeMessage    = "Bienvenidos a  FIUBADOS: "
	ObjectiveMessage  = "El objetivo de este juego es introducir a los alumnos que ingresan a la facultad en diversas cuestiones administrativas de su funcionamiento, siendo en el futuro muy útil para el desarrollo de su carrera profesional."
	MainMenuOptions   = "(1) Jugar  (2) Ayuda  (3) Salir"

	HelpMessage = "Este juego consiste en partidas 1vs1 en las que dos jugadores responden varias preguntas de opción múltiple. Cada jugador contestará el número de la opción que considere correcta en cada pregunta. El jugador que responda de forma correcta aumenta su puntuación. Si ambos jugadores responden correctamente una pregunta, el primero que haya respondido se lleva puntos adicionales. Asímismo, si solo uno de los jugadores responde correctamente, se lleva puntos adicionales. El jugador que responda incorrectamente no suma ningun punto."

	HelpMenuOptions = "(1) Volver al Menu principal"

	CloseConnectionCommand = "STOP"
	Success                = "OK"
	ReadyToPlay            = "LISTO"

	DisconnectAndExitMessage = "SERVER DISCONNECTED. Client exiting..."
	ExitMessage              = "Client exiting..."

	SearchingMatchMessage = "Armando una partida: Buscando jugadores..."

	PlayOption = "1"
	HelpOption = "2"
	ExitOption = "3"

	MatchingPlayerMessage = "Está jugando contra el jugador: "
	ReadyToPlayMessage    = ". Introduzca LISTO cuando esté listo para jugar"

	QuestionMessage = "Pregunta "
	ColonMessage    = ": "
	FirstOption     = " (1)"
	SecondOption    = " (2)"
	ThirdOption     = " (3)"

	CorrectAnswerMessage               = "Respuesta correcta! "
	WasFirstToAnswerMessage            = "Respondiste de primero! Recibes 3 puntos adicionales. "
	WasSecondToAnswerMessage           = "Respondiste de segundo! Recibes solo 1 punto. "
	OpponentAnsweredIncorrectlyMessage = "Tu contrincante respondió incorrectamente, te llevas 3 puntos adicionales. "
	IncorrectAnswerMessage             = "Respuesta incorrecta! "
	WhichWasCorrectAnswerMessage       = "La respuesta correcta era: ("

	PlayerMessage                  = "Jugador "
	WinnerMessage                  = " ha ganado! ¡Gracias por jugar a FIUBADOS! Pulsa cualquier tecla para salir"
	TieMessage                     = "¡Juego empatado! ¡Gracias por jugar a FIUBADOS! Pulsa cualquier tecla para salir"
	OtherPlayerDisconnectedMessage = "El otro jugador se desconectó ¡Ganaste, gracias por jugar a FIUBADOS! Pulsa cualquier tecla para salir"

	ColorCyan  = "\033[96m"
	ColorReset = "\033[0m"
	ColorGreen = "\033[92m"
)
