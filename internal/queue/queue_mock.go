package queue

import (
	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockQueue struct {
	mock.Mock
	jobs chan models.EmailRequest
}

func (q *MockQueue) Enqueue(job models.EmailRequest) error {
	args := q.Called(job)
	return args.Error(0)
}

func (q *MockQueue) Jobs() <-chan models.EmailRequest {
	return q.jobs
}
