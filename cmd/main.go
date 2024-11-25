package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"myproject/internal"
	"myproject/internal/mysql"
	"myproject/internal/nginx"
)

func main() {
	server, username, password := internal.InputData()

	// Create a single SSH connection
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	err = mysql.CheckMySQL(client)
	if err != nil {
		log.Printf("MySQL task failed: %v", err)
	}

	err = nginx.CheckNginx(client)
	if err != nil {
		log.Printf("Nginx task failed: %v", err)
	}
}
