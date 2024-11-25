package internal

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func CheckNginx(client *ssh.Client) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session for Nginx: %v", err)
	}
	defer session.Close()

	cmd := "grep 'error' /var/log/nginx/error.log"
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return fmt.Errorf("failed to check Nginx logs: %v, output: %s", err, string(output))
	}

	log.Printf("Nginx log errors:\n%s", string(output))
	return nil
}
