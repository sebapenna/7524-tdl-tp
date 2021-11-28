package logger

import (
	"fmt"
	"os"
)

func LogError(e error) {
	_, _ = fmt.Fprintln(os.Stderr, e)
}

func LogErrorMessage(msg string) {
	_, _ = fmt.Fprintln(os.Stderr, msg)
}

func LogInfo(a ...interface{}) {
	fmt.Println(a...)
}

func PrintMessageReceived(msg string) {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"
	fmt.Println(string(colorCyan), "->: "+msg, string(colorReset))
}
