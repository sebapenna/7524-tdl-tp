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
