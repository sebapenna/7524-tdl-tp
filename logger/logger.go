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
	} else if strings.Contains(msg, common.WinnerMessage) || strings.Contains(msg, common.OtherPlayerDisconnectedMessage) {
		fmt.Println(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
		fmt.Println(string(common.ColorYellow), common.AsciWinnerMessage, string(common.ColorReset))
	} else if strings.Contains(msg, common.TieMessage) {
		fmt.Println(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
		fmt.Println(string(common.ColorYellow), common.AsciTieMessage, string(common.ColorReset))
	} else {
		fmt.Println(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
	}

}
