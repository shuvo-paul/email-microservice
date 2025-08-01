package service

import (
	"testing"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/shuvo-paul/email-microservice/internal/queue"
	"github.com/stretchr/testify/assert"
)

func TestEmailService_Success(t *testing.T) {
	q := new(queue.MockQueue)
	s := NewEmailService(q)

	email := models.EmailRequest{
		To:      "test@example.org",
		Subject: "Test subject",
		Body:    "Test body",
	}

	q.On("Enqueue", email).Return(nil)
	err := s.Send(email)

	assert.NoError(t, err)
	q.AssertExpectations(t)
}

func TestEmailService_Send_QueueError(t *testing.T) {
	q := new(queue.MockQueue)
	s := NewEmailService(q)

	email := models.EmailRequest{
		To:      "test@example.org",
		Subject: "Test subject",
		Body:    "Test body",
	}

	q.On("Enqueue", email).Return(queue.ErrQueueFull)
	err := s.Send(email)

	assert.Equal(t, queue.ErrQueueFull, err)
	q.AssertExpectations(t)
}
