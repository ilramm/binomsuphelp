package mysql

import (
	"golang.org/x/crypto/ssh"
	"log"
	"myproject/internal"
	"strings"
)

func MysqlErrors() {
	server, username, password := internal.InputData()

	// Configure SSH connection
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the server
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()
	// First session: List files in /var/log/mysql
	sessionDir, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer sessionDir.Close()

	// List files in /var/log/mysql directory
	command1 := "ls /var/log/mysql"
	outputDir, err := sessionDir.CombinedOutput(command1)
	if err != nil {
		log.Fatalf("Failed to execute command: %v\nOutput: %s", err, string(outputDir))
	}

	// Debug: Print directory listing
	log.Printf("Directory listing output:\n%s", string(outputDir))

	// Parse the latest log file
	latestLog, err := GetLatestLogFile(string(outputDir))
	if err != nil {
		log.Fatalf("Failed to get latest log: %v", err)
	}

	// Debug: Print the latest log file name
	log.Printf("Latest log file identified: %s", latestLog)

	// Second session: Grep for errors in the latest log
	sessionLogRead, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer sessionLogRead.Close()

	// Use full path for the log file
	command2 := "grep 'ERROR' /var/log/mysql/" + latestLog
	outputLog, err := sessionLogRead.CombinedOutput(command2)
	if err != nil {
		log.Fatalf("Failed to execute command: %v\nOutput: %s", err, string(outputLog))
	}
	// Process log lines for errors
	logLines := strings.Split(string(outputLog), "\n")
	for _, line := range logLines {
		if strings.Contains(line, "ERROR") {
			log.Println(line)
		}
	}
}
