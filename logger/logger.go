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

func LogErrorMessage(msg ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, msg...)
}

func LogInfo(msg ...interface{}) {
	fmt.Println(msg...)
}

func PrintMessageReceived(msg string) {

	if strings.HasPrefix(common.AskForNameMessage, msg) {
		LogInfo(string(common.ColorCyan), common.ServerArrow+common.AsciAskForNameMessage, string(common.ColorReset))

	} else if strings.Contains(msg, common.WinnerMessage) || strings.Contains(msg, common.OtherPlayerDisconnectedMessage) {
		LogInfo(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
		LogInfo(string(common.ColorYellow), common.AsciWinnerMessage, string(common.ColorReset))

	} else if strings.Contains(msg, common.TieMessage) {
		LogInfo(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
		LogInfo(string(common.ColorYellow), common.AsciTieMessage, string(common.ColorReset))

	} else {
		LogInfo(string(common.ColorCyan), common.ServerArrow+msg, string(common.ColorReset))
	}

}
