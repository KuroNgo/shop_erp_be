package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type EmailData struct {
	Code      string
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

	fmt.Println("Am parsing templates...")

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(data *EmailData, emailTo string, templateName string) error {
	var body bytes.Buffer

	templated, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	err = templated.ExecuteTemplate(&body, templateName, &data)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", Mailer1)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", data.Subject)

	//m.SetAddressHeader(subject_const.Bcc, subject_const.BCCAdmin1, subject_const.Admin)
	//m.SetAddressHeader(subject_const.Bcc, subject_const.BCCAdmin2, subject_const.Admin)
	m.SetAddressHeader(Bcc, BCCAdmin3, Admin)

	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", body.String())

	d := gomail.NewDialer(SMTP_Host, SMTP_PORT, Mailer1, Password1)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err = d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
