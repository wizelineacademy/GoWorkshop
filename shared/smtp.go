package shared

import (
	"net/smtp"
)

// SendEmail helper
func SendEmail(to string, body string) error {
	auth := smtp.PlainAuth("",
		AppConfig.SMTPUser,
		AppConfig.SMTPPass,
		AppConfig.SMTPHost,
	)

	return smtp.SendMail(AppConfig.SMTPHost+":"+AppConfig.SMTPPort, auth, AppConfig.SMTPUser, []string{to}, []byte(body))
}
