// Package queue
package queue

import (
	"errors"

	"github.com/shuvo-paul/email-microservice/internal/models"
)

var ErrQueueFull = errors.New("queue full")

type Queue interface {
	Enqueue(models.EmailRequest) error
	Jobs() <-chan models.EmailRequest
}

var _ Queue = (*InMemoryQueue)(nil)

type InMemoryQueue struct {
	jobs chan models.EmailRequest
}

func NewQueue(size int) *InMemoryQueue {
	return &InMemoryQueue{
		jobs: make(chan models.EmailRequest, size),
	}
}

func (q *InMemoryQueue) Enqueue(emailRequest models.EmailRequest) error {
	select {
	case q.jobs <- emailRequest:
		return nil
	default:
		return ErrQueueFull
	}
}

func (q *InMemoryQueue) Jobs() <-chan models.EmailRequest {
	return q.jobs
}
