package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strconv"

	"github.com/karlozz157/storicard/src/domain/entity"
	"github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/domain/ports/service"
	"github.com/karlozz157/storicard/src/utils"
)

var logger = utils.GetLogger()

const (
	subject  = "Estado de Cuenta | Stori"
	htmlPath = "/templates/summary.html"
)

type MailerNotification struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewMailerNotification() service.INotificationService {
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	return &MailerNotification{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     port,
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}
}

func (s *MailerNotification) Notify(summary *entity.Summary) error {
	template, err := s.getTemplate(summary)
	if err != nil {
		return errors.ErrInternal
	}

	message := []byte(
		fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
			summary.Email,
			subject,
			template))

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", s.Host, s.Port), auth, s.Username, []string{summary.Email}, message)

	if err != nil {
		logger.Errorw("sending email", "err", err)
		return errors.ErrInternal
	}

	return err
}

func (s *MailerNotification) getTemplate(summary *entity.Summary) (string, error) {
	mainPath, _ := os.Getwd()
	templatePath := fmt.Sprintf("%s%s", mainPath, htmlPath)

	file, err := os.Open(templatePath)

	if err != nil {
		logger.Errorw("opening template file", "err", err)
		return "", err
	}
	defer file.Close()

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		logger.Errorw("parsing template template", "err", err)
		return "", err
	}

	var body bytes.Buffer
	err = template.Execute(&body, summary)
	if err != nil {
		logger.Errorw("rendering template", "err", err)
		return "", err
	}

	return body.String(), nil
}
