// Package worker
package worker

import (
	"log"
	"sync"

	"github.com/shuvo-paul/email-microservice/internal/mailer"
	"github.com/shuvo-paul/email-microservice/internal/queue"
)

type WorkerPool struct {
	queue  queue.Queue
	sender mailer.Sender
	wg     sync.WaitGroup
}

func NewWorkerPool(q queue.Queue, s mailer.Sender) *WorkerPool {
	return &WorkerPool{
		queue:  q,
		sender: s,
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
	for job := range p.queue.Jobs() {
		if err := p.sender.Send(job.To, job.Subject, job.Body); err != nil {
			// TODO: add to retry
			log.Printf("%v", err)
		}
	}
}

func (p *WorkerPool) ShutDown() {
	p.queue.Close()
	p.wg.Wait()
	log.Println("WorkerPool shut down gracefully!")
}
