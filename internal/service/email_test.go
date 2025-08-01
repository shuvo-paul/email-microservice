package service

import (
	"testing"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/shuvo-paul/email-microservice/internal/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockQueue struct {
	mock.Mock
	jobs chan models.EmailRequest
}

func (q *mockQueue) Enqueue(job models.EmailRequest) error {
	args := q.Called(job)
	return args.Error(0)
}

func (q *mockQueue) Jobs() <-chan models.EmailRequest {
	return q.jobs
}

func TestEmailService_Success(t *testing.T) {
	q := new(mockQueue)
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
	q := new(mockQueue)
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
