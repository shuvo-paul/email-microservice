# 🏗️ Architecture Document
## Email Notification Microservice

---

## 📌 1. Overview

This document outlines the architecture for the Email Notification Microservice including its major components, data flow, interactions, and deployment strategies.

---

## 🧱 2. Components

- **API Layer:** HTTP interface to accept email job requests and return responses.
- **Job Queue:** In-memory queue (channel) to buffer jobs for processing.
- **Worker Pool:** Goroutines processing email jobs in the background.
- **Retry Manager:** Handles failed job retries using exponential backoff.
- **Health Check Handler:** Monitors uptime and worker/queue status.
- **Logging Module:** Logs all job lifecycle events in structured format.
- **SMTP Client:** Connects to mail server for delivery.
- **Config Loader:** Loads env values and exposes application config.

---

## 🔄 3. Data Flow

1. API receives email job via `POST /send`
2. Validates payload
3. Enqueues job into channel
4. Worker pulls from channel, formats email
5. Connects to SMTP and attempts send
6. Logs success or hands off to Retry Manager
7. Retry Manager retries failed jobs up to `MAX_RETRIES`

---

## 🔁 4. Worker Pool Model

```go
jobs := make(chan EmailJob, 100)
for i := 0; i < WORKER_COUNT; i++ {
    go Worker(jobs)
}
```

Each worker reads from the channel. If sending fails, the job is handed to a retry channel.

---

## 🚦 5. Deployment Architecture

### Local Dev
- Run with `go run cmd/main.go`
- Uses `.env` for SMTP config

### Docker
- One container for app
- Exposes port `8080`
- Config from env or Docker Secrets

### Kubernetes
- `Deployment` + `Service`
- ConfigMaps for env
- Horizontal Pod Autoscaler (future)

---

## 🔧 6. Configuration

| Variable     | Purpose                  |
|--------------|---------------------------|
| SMTP_HOST    | SMTP server hostname      |
| SMTP_PORT    | Port (usually 587)        |
| SMTP_USER    | Auth username             |
| SMTP_PASS    | Auth password             |
| WORKER_COUNT | Number of goroutines      |
| MAX_RETRIES  | Max retry attempts        |
| QUEUE_SIZE   | Channel buffer size       |

---

## 📏 7. Observability

- `/health`: JSON status endpoint
- Future: `/metrics` Prometheus-compatible
- Logs: JSON to stdout (Zap/Zerolog)

---

## 🧩 8. Extensibility

- Queue: Replace channel with Redis, NATS, or RabbitMQ
- Email delivery: Swap SMTP with SendGrid/Mailgun SDK
- Auth: Add API key/JWT support
- Persistent Storage (if needed): PostgreSQL, SQLite, or Redis

---

## ✅ 9. Summary

This architecture enables a simple, fast, and reliable way to enqueue and send emails in background without blocking the main request/response flow. It can be scaled horizontally and adapted for larger workloads.

---
