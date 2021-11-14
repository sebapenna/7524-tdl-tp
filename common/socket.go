package common

import (
	"bufio"
	"fmt"
	"net"
)

func Receive(c net.Conn) (string, error) {
	return bufio.NewReader(c).ReadString('\n')
}

func Send(c net.Conn, text string) {
	fmt.Fprintf(c, text+"\n")
}
