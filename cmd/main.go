package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Read connection details from the user
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

	// Create the SSH client configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Not secure, replace in production.
	}

	// Establish the SSH connection
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	fmt.Println("SSH connection established")

	// Run a remote command
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Execute a command on the remote server
	fmt.Print("Enter a command to run on the remote server: ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	fmt.Printf("Command output:\n%s\n", output)
}
