package queue

import (
	"testing"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue_Success(t *testing.T) {
	q := NewQueue(2)

	job := models.EmailRequest{
		To:      "test@example.org",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	err := q.Enqueue(job)

	assert.NoError(t, err)

	queued := <-q.jobs
	assert.Equal(t, job, queued)
}

func TestQueue_Enqueue_Overflow(t *testing.T) {
	q := NewQueue(1)

	job := models.EmailRequest{
		To:      "test@example.org",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	q.Enqueue(job)
	err := q.Enqueue(job)

	assert.Error(t, err)
	assert.Equal(t, err, ErrQueueFull)
}
