package utils

import (
	"bytes"
	"crypto/tls"
	"github.com/NicholasLiem/Email_Confirmation_Service_Go/schema"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *schema.User, data *EmailData) {

	from := os.Getenv("EMAIL_FROM")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")
	to := user.Email
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	smtpPortInt, err := strconv.ParseUint(smtpPort, 10, 64)
	if err != nil {
		return
	}

	var body bytes.Buffer

	parseTemplateDir, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse parseTemplateDir", err)
	}

	err = parseTemplateDir.ExecuteTemplate(&body, "verificationCode.html", &data)
	if err != nil {
		return
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, int(smtpPortInt), smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
	}

}
