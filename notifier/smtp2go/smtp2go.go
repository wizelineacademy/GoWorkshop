package smtp2go

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

// SendEmail : sends email using smtp2go testing SMTP service
func SendEmail(to string, subject string, body string) error {
	pass, err := base64.StdEncoding.DecodeString("WlVjdGJ2VGxBbDBM")
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("",
		"alexander.plutov@gmail.com",
		string(pass),
		"mail.smtp2go.com",
	)

	header := make(map[string]string)
	header["From"] = "alexander.plutov@gmail.com"
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

	return smtp.SendMail("mail.smtp2go.com:587", auth, "alexander.plutov@gmail.com", []string{to}, []byte(message))
}
