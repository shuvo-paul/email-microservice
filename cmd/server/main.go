package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shuvo-paul/email-microservice/internal/config"
	"github.com/shuvo-paul/email-microservice/internal/handlers"
	"github.com/shuvo-paul/email-microservice/internal/mailer"
	"github.com/shuvo-paul/email-microservice/internal/queue"
	"github.com/shuvo-paul/email-microservice/internal/service"
	"github.com/shuvo-paul/email-microservice/internal/worker"
)

func main() {
	cfg := config.Load()
	queue := queue.NewQueue(10)
	emailService := service.NewEmailService(queue)
	emailHandler := handlers.NewEmailHandler(emailService)
	sender := mailer.NewSMTPSender(
		cfg.SMTPConfig.Host,
		cfg.SMTPConfig.Port,
		cfg.SMTPConfig.Username,
		cfg.SMTPConfig.Password,
		cfg.SMTPConfig.From,
	)
	wp := worker.NewWorkerPool(queue, &sender)
	wp.Start(5)

	http.HandleFunc("GET /health", handlers.HealthHandler)
	http.HandleFunc("/send", emailHandler.ServeHTTP)

	port := fmt.Sprintf(":%d", cfg.ServerConfig.Port)
	log.Println("Server starting on :" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
