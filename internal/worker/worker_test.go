package worker

import (
	"testing"
	"time"

	"github.com/shuvo-paul/email-microservice/internal/mailer"
	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/shuvo-paul/email-microservice/internal/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewWorkerPool(t *testing.T) {
	mockQueue := queue.NewMockQueue(make(chan models.EmailRequest, 10))
	mockSender := &mailer.MockSender{}

	pool := NewWorkerPool(mockQueue, mockSender)

	assert.NotNil(t, pool)
	assert.Equal(t, pool.queue, mockQueue)
	assert.Equal(t, pool.sender, mockSender)
}

func TestWorkerPool_Start_Success(t *testing.T) {
	jobChan := make(chan models.EmailRequest, 10)
	mockQueue := queue.NewMockQueue(jobChan)
	mockSender := mailer.NewMockSender()

	job := models.EmailRequest{
		To:      "user@example.org",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	mockSender.On("Send", job.To, job.Subject, job.Body).Return(nil).Times(2)
	mockQueue.On("Close").Return()

	pool := NewWorkerPool(mockQueue, mockSender)

	pool.Start(2)

	for range 2 {
		jobChan <- job
	}
	pool.ShutDown()

	mockSender.AssertExpectations(t)
	mockQueue.AssertExpectations(t)
}

func TestWorkerPool_ShutDown(t *testing.T) {
	jobChan := make(chan models.EmailRequest, 10)
	mockQueue := queue.NewMockQueue(jobChan)
	mockSender := mailer.NewMockSender()

	mockQueue.On("Close").Return().Once()

	pool := NewWorkerPool(mockQueue, mockSender)
	pool.Start(2)
	pool.ShutDown()

	mockQueue.AssertExpectations(t)

	_, ok := <-jobChan
	assert.False(t, ok)
}

func TestWorkerPool_Concurrency(t *testing.T) {
	workerCount := 3

	jobChan := make(chan models.EmailRequest, 10)
	mockQueue := queue.NewMockQueue(jobChan)
	mockSender := mailer.NewMockSender()
	pool := NewWorkerPool(mockQueue, mockSender)

	readyChan := make(chan struct{}, workerCount)
	blockChan := make(chan struct{})

	mockSender.On("Send", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		readyChan <- struct{}{}
		<-blockChan
	}).Return(nil).Times(workerCount)

	mockQueue.On("Close").Return()

	pool.Start(workerCount)

	for range workerCount {
		jobChan <- models.EmailRequest{To: "test@example.org"}
	}

	for i := range workerCount {
		select {
		case <-readyChan:
		case <-time.After(1 * time.Second):
			t.Fatalf("Time out waiting for woker %d to start. Concurrency check failed", i)
		}
	}

	close(blockChan)
	pool.ShutDown()

	mockSender.AssertExpectations(t)
}
