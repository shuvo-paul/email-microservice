// Package mailer provides a simple interface for sending emails using SMTP
package mailer

import (
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

type Sender interface {
	Send(to, subject, body string) error
}

type SMTPSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func (s *SMTPSender) Send(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	var msg []byte
	msg = fmt.Appendf(msg, "From: %s\r\nTo: %s\r\nSubject: %s\r\nBody: %s", s.from, to, subject, body)

	return smtp.SendMail(addr, auth, s.from, []string{to}, msg)
}

type Client struct {
	sender Sender
}

func NewClient(host string, port int, username string, password string, from string) *Client {
	return &Client{
		sender: &SMTPSender{
			host:     host,
			port:     port,
			username: username,
			password: password,
			from:     from,
		},
	}
}

func (d *Client) Send(to, subject, body string) error {
	if _, err := mail.ParseAddress(to); err != nil {
		return fmt.Errorf("invalid recipient address: %w", err)
	}
	cleanSubject, err := sanitizeSubject(subject)
	if err != nil {
		return fmt.Errorf("invalid subject: %w", err)
	}

	return d.sender.Send(to, cleanSubject, body)
}

func sanitizeSubject(subject string) (string, error) {
	if strings.Contains(subject, "\n") || strings.Contains(subject, "\r") {
		return "", errors.New("subject cannot contain newlines")
	}

	if len(subject) > 200 {
		return "", errors.New("subject too long (max 200 characters)")
	}
	cleaned := strings.Map(func(r rune) rune {
		if r < 32 && r != 9 {
			return -1
		}
		return r
	}, subject)
	return cleaned, nil
}
