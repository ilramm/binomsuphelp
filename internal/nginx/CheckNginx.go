package nginx

import (
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
)

func CheckNginx(client *ssh.Client) error {

	// первая сессия
	sessionDir, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer sessionDir.Close()

	// листим логи
	command1 := "ls /var/log/nginx"
	outputDir, err := sessionDir.CombinedOutput(command1)
	if err != nil {
		log.Fatalf("Failed to execute command: %v\nOutput: %s", err, string(outputDir))
	}

	// дебах
	log.Printf("Directory listing output:\n%s", string(outputDir))

	// парсим лог
	latestLog, err := GetLatestLogFile(string(outputDir))
	if err != nil {
		log.Fatalf("Failed to get latest log: %v", err)
	}

	// дебах
	log.Printf("Latest log file identified: %s", latestLog)

	// греп логов
	sessionLogRead, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer sessionLogRead.Close()

	// полный путь к логм
	command2 := "grep 'ERROR' /var/log/nginx/" + latestLog
	outputLog, err := sessionLogRead.CombinedOutput(command2)
	if err != nil {
		log.Fatalf("Failed to execute command: %v\nOutput: %s", err, string(outputLog))
	}
	// ищем ошибки
	logLines := strings.Split(string(outputLog), "\n")
	for _, line := range logLines {
		if strings.Contains(line, "ERROR") {
			log.Println(line)
		}
	}
	return nil
}
