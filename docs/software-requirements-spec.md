# 📄 Software Requirements Specification (SRS)
## Email Notification Microservice

---

## 📌 1. Introduction

### 1.1 Purpose
This document defines the functional and non-functional requirements for the Email Notification Microservice. The service is responsible for accepting and processing email requests asynchronously, ensuring reliable delivery using SMTP or a third-party email provider.

### 1.2 Scope
- The system will expose an HTTP API to enqueue email jobs.
- Jobs will be processed by a pool of background workers.
- Failed jobs will be retried using a backoff strategy.
- The service is self-contained and will serve as a utility microservice in a distributed system.

### 1.3 Intended Audience
- Developers and DevOps engineers
- QA testers
- System architects
- API consumers (internal/backend services)

### 1.4 Definitions and Acronyms
| Term | Definition |
|------|------------|
| SMTP | Simple Mail Transfer Protocol |
| API  | Application Programming Interface |
| NFR  | Non-Functional Requirements |
| FR   | Functional Requirements |

---

## 🚀 2. Overall Description

### 2.1 Product Perspective
This microservice is designed to be a reusable component within a microservice architecture. It is stateless, horizontally scalable, and decoupled from consumers through a RESTful API.

### 2.2 Product Functions
- Accept email job requests via a REST endpoint.
- Validate and sanitize input.
- Enqueue jobs for background processing.
- Send emails using SMTP or a configured provider.
- Handle failures and retry jobs up to a configured threshold.
- Expose health and readiness endpoints for orchestration.

### 2.3 User Classes and Characteristics
| User Type     | Description |
|---------------|-------------|
| API Consumer  | Sends email job requests via HTTP |
| Admin/DevOps  | Monitors logs, health endpoints, and system performance |
| Developer     | Extends or maintains the codebase |

---

## 🧩 3. Functional Requirements (FRs)

| ID | Requirement |
|----|-------------|
| FR1 | The system shall expose a `POST /send` endpoint to accept email job requests. |
| FR2 | The system shall validate required fields: `to`, `subject`, and `body`. |
| FR3 | The system shall enqueue valid jobs into an in-memory queue. |
| FR4 | The system shall run a pool of N workers to process jobs concurrently. |
| FR5 | The system shall send emails using SMTP or a configured provider. |
| FR6 | The system shall retry failed jobs up to `MAX_RETRIES` with exponential backoff. |
| FR7 | The system shall log successful and failed deliveries. |
| FR8 | The system shall expose a `GET /health` endpoint for health checks. |
| FR9 | The system shall reject requests if the queue is full with a 503 status. |
| FR10 | The system shall load configuration from environment variables or a `.env` file. |

---

## 📐 4. Non-Functional Requirements (NFRs)

| ID | Requirement |
|----|-------------|
| NFR1 | The system shall process 1000 jobs per second under load. |
| NFR2 | The system shall achieve 99.9% uptime. |
| NFR3 | The system shall support horizontal scaling of workers. |
| NFR4 | The system shall be containerized using Docker. |
| NFR5 | The system shall deliver emails within 5 seconds of request under nominal load. |
| NFR6 | The system shall provide structured JSON logs for observability. |
| NFR7 | The system shall gracefully shut down, completing in-flight jobs. |
| NFR8 | The system shall handle SMTP timeouts and connection errors robustly. |

---

## 🛑 5. Constraints

- SMTP credentials must be securely provided via environment variables.
- No persistence layer (e.g. database) will be used initially.
- Email body size must not exceed 50KB.
- Only one destination address (`to`) per email for MVP.

---

## 🔗 6. External Interfaces

### 6.1 API Interface
- `POST /send` – Enqueue an email job  
- `GET /health` – Health check  

### 6.2 Environment Variables
| Variable | Description |
|----------|-------------|
| SMTP_HOST | SMTP server hostname |
| SMTP_PORT | SMTP port |
| SMTP_USER | SMTP username |
| SMTP_PASS | SMTP password |
| WORKER_COUNT | Number of concurrent workers |
| MAX_RETRIES | Max retry attempts for a failed job |
| QUEUE_SIZE | Max number of jobs in the queue |

---

## 📦 7. Assumptions and Dependencies

- Email jobs are ephemeral and do not require durable storage for now.
- All consumers are trusted internal systems (no auth required initially).
- SMTP server is available and supports TLS and authentication.

---

## 📅 8. Future Considerations

- Support for batch emails and CC/BCC
- Email templates and localization
- OpenAPI documentation and SDK generation
- Integration with Mailgun, SendGrid, etc.
- Persistent job storage and replay (PostgreSQL or Redis)
- Rate limiting and throttling controls

---
