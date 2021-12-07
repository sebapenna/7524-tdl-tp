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

	if strings.HasPrefix(common.AskForNameMessage, msg) {
		fmt.Println(string(common.ColorCyan), common.ServerArrow+common.AsciAskForNameMessage, string(common.ColorReset))
		//} else if strings.HasPrefix(common.PlayerMessage, msg) {
		//	fmt.Println(string(common.ColorCyan), common.ServerArrow+msg+hola, string(common.ColorReset))
	} else {
		fmt.Println(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
	}

}
