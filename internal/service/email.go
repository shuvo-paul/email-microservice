// Package service contains the core business logic for processing and sending email logic
package service

import (
	"errors"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/shuvo-paul/email-microservice/internal/queue"
)

var ErrSendingEmail = errors.New("email sending failed")

type Sender interface {
	Send(models.EmailRequest) error
}
type EmailService struct {
	queue queue.Queue
}

func NewEmailService(q queue.Queue) *EmailService {
	return &EmailService{
		queue: q,
	}
}

func (e *EmailService) Send(email models.EmailRequest) error {
	err := e.queue.Enqueue(email)
	if err != nil {
		return err
	}
	return nil
}
