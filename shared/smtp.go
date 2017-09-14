package shared

import (
	"net/smtp"
)

// SendEmail helper
func SendEmail(to string, body string) error {
	auth := smtp.PlainAuth("",
		AppConfig.SmtpUser,
		AppConfig.SmtpPass,
		AppConfig.SmtpHost,
	)

	return smtp.SendMail(AppConfig.SmtpHost+":"+AppConfig.SmtpPort, auth, AppConfig.SmtpUser, []string{to}, []byte(body))
}
