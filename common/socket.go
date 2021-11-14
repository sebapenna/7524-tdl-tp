package common

import (
	"bufio"
	"fmt"
	"io"
)

func Receive(c io.Reader) (string, error) {
	return bufio.NewReader(c).ReadString('\n')
}

func Send(c io.Writer, text string) {
	fmt.Fprintf(c, text+"\n")
}
