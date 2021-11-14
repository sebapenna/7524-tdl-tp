package logger

import (
	"fmt"
	"os"
)

func LogError(e error) {
	_, _ = fmt.Fprint(os.Stderr, e)
}

func LogErrorMessage(msg string) {
	_, _ = fmt.Fprint(os.Stderr, msg)
}
