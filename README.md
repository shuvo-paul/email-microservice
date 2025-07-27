# 📬 Email Notification Microservice

A lightweight, concurrent, and fault-tolerant microservice written in Go that accepts email requests and delivers them via SMTP or a third-party email API.

## 🔧 Features
- REST API for enqueuing emails
- Worker pool for background processing
- Retry system with exponential backoff
- Configurable via `.env` or Docker
- Health check and future metrics support

## 🚀 Quick Start
```bash
cp .env.example .env
go run cmd/main.go
# or
docker-compose up --build
```

## 📚 Documentation
- [Software Requirements Specification (SRS)](docs/software-requirements-spec.md)
- [Functional Spec](docs/functional-spec.md)
- [Non-Functional Requirements](docs/non-functional-requirements.md)
- [Architecture](docs/architecture.md)
- [Deployment](docs/deployment.md)
- [Process](docs/process.md)
- [Strategy Roadmap](docs/strategy-roadmap.md)
- [Technology Roadmap](docs/technology-roadmap.md)
- [UX](docs/ux.md)
- [Release Roadmap](docs/release-roadmap.md)
- [Contributing Guide](docs/CONTRIBUTING.md)
- [OpenAPI Spec](docs/api/openapi.yaml)

