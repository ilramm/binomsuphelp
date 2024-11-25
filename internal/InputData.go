package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputData() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter server IP (e.g., 192.168.1.1:22): ")
	server, _ := reader.ReadString('\n')
	server = strings.TrimSpace(server)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	return server, username, password
}
