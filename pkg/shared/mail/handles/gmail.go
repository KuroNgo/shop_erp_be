package handles

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"path/filepath"
	"shop_erp_mono/pkg/shared/mail/constant"
)

type EmailData struct {
	Code        string
	FullName    string
	Subject     string
	HREmail     string
	EmployeeID  string
	LeaveType   string
	ProductList []string
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

	templated, err := ParseTemplateDir("pkg/shared/mail/templates")
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	err = templated.ExecuteTemplate(&body, templateName, data)
	if err != nil {
		return fmt.Errorf("could not execute template: %v", err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", constant.Mailer1)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", data.Subject)

	m.SetAddressHeader(constant.Bcc, constant.BCCAdmin3, constant.Admin)

	m.SetBody("text/html", body.String())
	//m.AddAlternative("text/plain", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(constant.SMTP_Host, constant.SMTP_PORT, constant.Mailer1, constant.Password1)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err = d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
