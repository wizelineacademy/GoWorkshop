package shared

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

func SendEmail(to string, subject string, body string) error {
	pass, _ := base64.StdEncoding.DecodeString(AppConfig.SMTPPass)
	auth := smtp.PlainAuth("",
		AppConfig.SMTPUser,
		string(pass),
		AppConfig.SMTPHost,
	)

	header := make(map[string]string)
	header["From"] = AppConfig.SMTPUser
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return smtp.SendMail(AppConfig.SMTPHost+":"+AppConfig.SMTPPort, auth, AppConfig.SMTPUser, []string{to}, []byte(message))
}
