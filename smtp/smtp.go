package smtpModule

import (
	"audio-stream-golang/config"
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(to string, subject string, body string, token string) error {
	EnvConfig := config.GetConfig()
	smtp_host := EnvConfig.Smtp_host
	smtp_port := EnvConfig.Smtp_port
	smtp_sender := EnvConfig.Smtp_sender
	message := fmt.Sprintf("Subkect: %s\r\n\r\n%s", subject, body+token)
	//recipient := []string{to}
	err := smtp.SendMail(smtp_host+":"+smtp_port, nil, smtp_sender, []string{to}, []byte(message))
	if err != nil {
		log.Println("error!", err)
	}

	return nil
}
