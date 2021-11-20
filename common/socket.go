package common

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Receive(c net.Conn) (string, error) {
	str, err := bufio.NewReader(c).ReadString('\n')
	return strings.TrimSpace(str), err
}

func Send(c net.Conn, text string) {
	fmt.Fprintf(c, text+"\n")
}
