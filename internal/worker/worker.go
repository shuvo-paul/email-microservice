// Package worker
package worker

import (
	"context"
	"log"
	"sync"

	"github.com/shuvo-paul/email-microservice/internal/mailer"
	"github.com/shuvo-paul/email-microservice/internal/queue"
)

type WorkerPool struct {
	queue  queue.Queue
	sender mailer.Sender
	wg     sync.WaitGroup
	once   sync.Once
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorkerPool(q queue.Queue, s mailer.Sender) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())

	return &WorkerPool{
		queue:  q,
		sender: s,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (p *WorkerPool) Start(n int) {
	p.wg.Add(n)
	for range n {
		go p.worker()
	}
}

func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for {
		select {
		case job, ok := <-p.queue.Jobs():
			if !ok {
				return
			}
			if err := p.sender.Send(job.To, job.Subject, job.Body); err != nil {
				// TODO: add to retry
				log.Printf("%v", err)
			}
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *WorkerPool) ShutDown() {
	p.cancel()
	p.wg.Wait()
	log.Println("WorkerPool shut down gracefully!")
}
