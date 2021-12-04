package logger

import (
	"fmt"
	"os"

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
	//colorCyan := "\033[96m"
	//colorReset := "\033[0m"
	fmt.Println(string(common.ColorCyan), "->: "+msg, string(common.ColorReset))
}
