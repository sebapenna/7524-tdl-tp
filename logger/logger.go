package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sebapenna/7524-tdl-tp/common"
)

func LogError(e error) {
	_, _ = fmt.Fprintln(os.Stderr, e)
}

func LogErrorMessage(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}

func LogInfo(a ...interface{}) {
	fmt.Println(a...)
}

func PrintMessageReceived(msg string) {

	multiple := `🤖->:`
	askForNameAsci := `

    █▀█ █▀█ █▀█   █▀▀ ▄▀█ █░█ █▀█ █▀█ ░   █ █▄░█ ▀█▀ █▀█ █▀█ █▀▄ █░█ █▀▀ █▀▀   █░█ █▄░█   █▄░█ █▀█ █▀▄▀█ █▄▄ █▀█ █▀▀
    █▀▀ █▄█ █▀▄   █▀░ █▀█ ▀▄▀ █▄█ █▀▄ █   █ █░▀█ ░█░ █▀▄ █▄█ █▄▀ █▄█ █▄▄ ██▄   █▄█ █░▀█   █░▀█ █▄█ █░▀░█ █▄█ █▀▄ ██▄

    █▀█ ▄▀█ █▀█ ▄▀█   ░░█ █░█ █▀▀ ▄▀█ █▀█ ▀
    █▀▀ █▀█ █▀▄ █▀█   █▄█ █▄█ █▄█ █▀█ █▀▄ ▄
        `
	if strings.HasPrefix("Por favor,introduce un nombre para jugar: ", msg) {

		fmt.Println(string(common.ColorCyan), multiple+askForNameAsci, string(common.ColorReset))
	} else {
		fmt.Println(string(common.ColorCyan), multiple+msg, string(common.ColorReset))
	}

}
