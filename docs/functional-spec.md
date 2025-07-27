# 📄 Functional Specification Document
## Email Notification Microservice

---

## 📌 1. Introduction

This document provides detailed specifications of all the core functionalities defined in the Software Requirements Specification (SRS). It focuses on inputs, outputs, expected behaviors, error handling, and edge cases for each feature.

---

## 🧩 2. Functional Use Cases

---

### 🔹 2.1 Enqueue Email Job (`POST /send`)

**Description:**  
Accepts an email job and adds it to the in-memory queue.

**Request:**
```json
{
  "to": "user@example.com",
  "subject": "Welcome to the Platform",
  "body": "<p>Thanks for joining!</p>",
  "is_html": true
}
```

**Field Specifications:**

| Field     | Type    | Required | Constraints                     |
|-----------|---------|----------|----------------------------------|
| `to`      | string  | ✅       | Valid email format               |
| `subject` | string  | ✅       | Max 255 characters               |
| `body`    | string  | ✅       | Max 50KB                         |
| `is_html` | boolean | ❌       | Defaults to `false`              |

**Behavior:**
- Validates all fields.
- If valid, enqueues the job into the worker queue.
- If invalid, responds with 400 Bad Request.

**Responses:**

| Status | Description                    |
|--------|--------------------------------|
| 200    | Email enqueued successfully    |
| 400    | Missing or invalid input       |
| 503    | Queue full or service degraded |

---

### 🔹 2.2 Process Job from Queue (Worker Logic)

**Description:**  
Each worker pulls jobs from the queue and sends the email.

**Steps:**
1. Worker dequeues an `EmailJob`.
2. Constructs email with appropriate content type (`text/plain` or `text/html`).
3. Connects to SMTP server securely.
4. Sends the email using SMTP protocol.
5. Logs success or failure.

**Success Condition:**  
- Email is sent with a `250 OK` SMTP response code.

**Failure Condition:**  
- SMTP failure, timeout, TLS error, or rejection.
- Job is then passed to the retry manager.

---

### 🔹 2.3 Retry Failed Jobs

**Description:**  
Automatically re-enqueues failed jobs using exponential backoff up to a maximum retry limit.

**RetryJob Fields:**

| Field       | Description                     |
|-------------|---------------------------------|
| `Attempts`  | Number of attempts made so far  |
| `MaxRetry`  | Configured retry threshold      |
| `Delay`     | Delay before next retry attempt |

**Backoff Strategy:**
```
delay = baseDelay * (2 ^ attempts)
```

**Behavior:**
- If a job fails, it is wrapped in a `RetryJob` and sent to the retry queue.
- Retries are delayed and reprocessed by a separate goroutine.
- When `Attempts >= MaxRetry`, the job is logged as permanently failed.

---

### 🔹 2.4 Health Check (`GET /health`)

**Description:**  
Provides a simple liveness and readiness probe to external services or orchestrators.

**Response:**
```json
{
  "status": "ok",
  "uptime": "16m34s",
  "queue_depth": 23,
  "workers_active": 5
}
```

**Status Codes:**

| Status | Description                      |
|--------|----------------------------------|
| 200    | Service is operational           |
| 503    | Workers are failing or queue full |

---

## 📊 3. Error Handling

### Validation Errors

| Field     | Problem                   | Response                     |
|-----------|---------------------------|------------------------------|
| `to`      | Invalid email format       | 400 + error message          |
| `subject` | Empty or too long          | 400                          |
| `body`    | Missing or exceeds size    | 400                          |

### System Errors

| Scenario                     | Response                 |
|------------------------------|--------------------------|
| SMTP server unreachable      | Retry and 500 logged     |
| SMTP authentication failure  | Retry and log as error   |
| Queue full                   | Respond with 503         |
| Worker crash                 | Recovered by pool manager|

---

## 🛑 4. Edge Cases

| Case                                             | Handling                                                      |
|--------------------------------------------------|---------------------------------------------------------------|
| `is_html` is missing                             | Defaults to `false`                                           |
| `body` is HTML but `is_html=false`               | Sent as plain text                                            |
| SMTP times out mid-transaction                   | Job retried with backoff                                      |
| Retry attempts exhausted                         | Job logged as failed permanently                              |
| Queue overflows during high load                 | API returns 503 and logs overload                             |
| Job succeeds but logs fail (e.g., disk full)     | Email delivery considered valid; warning logged               |

---

## 🔗 5. Integration Points

- **SMTP Server**: Authenticated connection over TLS (configurable via environment).
- **Logging System**: Structured logs using JSON format (Zap or Zerolog recommended).
- **Queue System**: In-memory (Go channels); pluggable for Redis or RabbitMQ in future.
- **Monitoring/Probes**: Exposed via `/health` and (optionally) `/metrics`.

---

## ✅ 6. Success Criteria

- 95% of emails are successfully delivered within 5 seconds during normal load.
- All failed jobs are retried as per configured backoff policy.
- System should gracefully handle burst loads up to the configured queue capacity.
- Logs are structured and provide traceable details for each job's lifecycle.

---
