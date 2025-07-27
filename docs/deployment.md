# 🚀 Deployment Guide
## Email Notification Microservice

---

## 📦 1. Deployment Methods

### 🧪 Local Development

```bash
cp .env.example .env
go run cmd/main.go
```

- Uses `.env` for configuration
- SMTP credentials and worker settings passed via env vars

### 🐳 Docker

**Dockerfile**
```dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o email-service cmd/main.go
CMD ["./email-service"]
```

**docker-compose.yml**
```yaml
version: "3.8"
services:
  email-service:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
```

**Run**
```bash
docker-compose up --build
```

---

## ☁ Kubernetes (Planned)

- Define `Deployment` and `Service`
- Use `ConfigMap` for envs
- Set readiness/liveness probes
- Use `HorizontalPodAutoscaler` for scaling

---

## 🔐 Secrets Management

- Local: `.env`
- Docker: `.env` or `docker secrets`
- Kubernetes: `Secrets` or Vault (recommended)

---

## 🧪 Pre-Deploy Checklist

- [ ] `.env` is configured with SMTP and worker values
- [ ] Health endpoint returns `200 OK`
- [ ] Retry queue is empty
- [ ] Worker logs show at least one successful delivery

---

## 📤 CI/CD (Planned)

1. Lint + Test
2. Build binary or Docker image
3. Tag with `git` version
4. Push to container registry
5. Deploy to staging/production

---

## ✅ Post-Deploy Verification

- Send test job
- Observe logs
- Verify email delivery
- Confirm retry behavior on simulated failure

---
