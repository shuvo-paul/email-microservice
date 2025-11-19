package queue

import (
	"sync"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockQueue struct {
	mock.Mock
	jobs      chan models.EmailRequest
	closeOnce sync.Once
}

func NewMockQueue(jobs chan models.EmailRequest) *MockQueue {
	return &MockQueue{
		Mock: mock.Mock{},
		jobs: jobs,
	}
}

func (q *MockQueue) Enqueue(job models.EmailRequest) error {
	args := q.Called(job)
	return args.Error(0)
}

func (q *MockQueue) Jobs() <-chan models.EmailRequest {
	return q.jobs
}

func (q *MockQueue) Close() {
	q.closeOnce.Do(func() {
		if q.jobs != nil {
			close(q.jobs)
		}
	})
	q.Called()
}
