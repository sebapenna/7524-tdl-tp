package logger

import (
	"fmt"
	"os"
)

func LogError(e error) {
	_, _ = fmt.Fprintln(os.Stderr, e)
}

func LogErrorMessage(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a)
}

func LogInfo(a ...interface{}) {
	fmt.Println(a...)
}

func PrintMessageReceived(msg string) {
	colorCyan := "\033[36m"
	colorReset := "\033[0m"
	fmt.Println(string(colorCyan), "->: "+msg, string(colorReset))
}
