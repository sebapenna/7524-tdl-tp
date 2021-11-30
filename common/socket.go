package common

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Receive(connectionSocket net.Conn) (string, error) {
	str, err := bufio.NewReader(connectionSocket).ReadString('\n')
	return strings.TrimSpace(str), err
}

func Send(connectionSocket net.Conn, text string) error {
	_, err := fmt.Fprintf(connectionSocket, text+"\n")
	return err
}
