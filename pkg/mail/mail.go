package mail

import (
	"api/internal/domain"
	"api/pkg/configuration"
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"

	"github.com/jordan-wright/email"
	"go.uber.org/zap"
)

type MailService struct {
	logger *zap.Logger
	e      *email.Email

	host     string
	port     int
	username string
	password string
}

func New(logger *zap.Logger, cfg *configuration.SMTP) *MailService {
	return &MailService{
		e:        email.NewEmail(),
		logger:   logger,
		host:     cfg.Host,
		port:     cfg.Port,
		username: cfg.Username,
		password: cfg.Password,
	}
}

type mailCodeData struct {
	Url string
	Key string
}

func (m *MailService) SendCode(to, subject, templatePath string, key string) error {
	from := "Dapso Team <" + m.username + ">"
	data := mailCodeData{
		Key: key,
	}
	return m.sendWithTemplate(from, to, subject, templatePath, data)
}

func (m *MailService) sendWithTemplate(from, to, subject, templatePath string, data interface{}) error {
	// load template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		m.logger.Error("mail error", zap.Error(err))
		return domain.ErrMail
	}

	// set data to the template
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		m.logger.Error("mail error", zap.Error(err))
		return domain.ErrMail
	}

	// setup mail body
	m.e.From = from
	m.e.To = []string{to}
	m.e.Subject = subject
	m.e.HTML = body.Bytes()

	addr := fmt.Sprintf("%s:%d", m.host, m.port)

	// send the message
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         m.host,
	}
	err = m.e.SendWithTLS(addr, LoginAuth(m.username, m.password), tlsConfig)
	if err != nil {
		m.logger.Error("mail error", zap.Error(err))
		return domain.ErrMail
	}

	return nil
}
