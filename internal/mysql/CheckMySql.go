package mysql

import (
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
)

func CheckMySQL(client *ssh.Client) error {

	sessionDir, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer sessionDir.Close()

	command1 := "ls /var/log/mysql"
	outputDir, err := sessionDir.CombinedOutput(command1)
	if err != nil {
		log.Fatalf("Failed to execute command: %v\nOutput: %s", err, string(outputDir))
	}

	log.Printf("Directory listing output:\n%s", string(outputDir))

	latestLog, err := GetLatestLogFile(string(outputDir))
	if err != nil {
		log.Fatalf("Failed to get latest log: %v", err)
	}

	log.Printf("Latest log file identified: %s", latestLog)

	sessionLogRead, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer sessionLogRead.Close()

	command2 := "grep 'ERROR' /var/log/mysql/" + latestLog
	outputLog, err := sessionLogRead.CombinedOutput(command2)
	if err != nil {
		log.Fatalf("Failed to execute command: %v\nOutput: %s", err, string(outputLog))
	}
	logLines := strings.Split(string(outputLog), "\n")
	for _, line := range logLines {
		if strings.Contains(line, "ERROR") {
			log.Println(line)
		}
	}
	return nil
}
