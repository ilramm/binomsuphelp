package internal

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func CheckMySQL(client *ssh.Client) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session for MySQL: %v", err)
	}
	defer session.Close()

	cmd := "grep 'ERROR' /var/log/mysql/mysql_error.log"
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return fmt.Errorf("failed to check MySQL logs: %v, output: %s", err, string(output))
	}

	log.Printf("MySQL log errors:\n%s", string(output))
	return nil
}
