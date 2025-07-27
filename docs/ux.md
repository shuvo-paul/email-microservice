# 🧩 UX Documentation
## Developer & API Consumer Experience – Email Notification Microservice

---

## 📌 1. Introduction

This document outlines user experience (UX) considerations for **API consumers** and **developers** interacting with the Email Notification Microservice. Though this service has no frontend interface, it prioritizes developer-facing consistency, simplicity, and reliability in its API behavior.

---

## ✨ 2. UX Goals

- Simple and intuitive REST API
- Predictable responses with clear feedback
- Minimal setup required to send emails
- Graceful handling of overload, failure, and malformed input
- Useful and structured error messages for debugging
- Idempotent and stateless endpoints

---

## 📬 3. API Response Design

### 3.1 Success Response (200 OK)

```json
{
  "message": "Email enqueued successfully"
}
```

### 3.2 Validation Error (400 Bad Request)

```json
{
  "error": "Missing required field: 'to'"
}
```

### 3.3 Queue Overload (503 Service Unavailable)

```json
{
  "error": "Email queue is full. Please retry later."
}
```

---

## ⚠ 4. Error Handling Strategy

- All errors return structured JSON with an `error` field.
- HTTP status codes reflect actual failure types (400, 500, 503).
- No HTML error pages or raw panic messages are exposed to the client.

---

## 📦 5. Developer Environment Setup

| Step              | UX Feature                             |
|-------------------|-----------------------------------------|
| `.env` File       | Simple config loading for quick setup   |
| Docker Support    | One-command bootstrapping               |
| Health Endpoint   | Fast check for service readiness        |
| Local Testing     | curl/Postman-compatible `POST /send`    |

---

## 🧪 6. Testability and DX

- The service supports local test runs without external dependencies.
- Mocks can be injected to test email sending logic.
- Unit tests cover enqueueing, validation, and retry logic.

---

## 🛡 7. Fail-Safe Behavior

| Scenario                      | Behavior                                              |
|-------------------------------|-------------------------------------------------------|
| SMTP failure                  | Job is retried silently; API caller not affected     |
| Invalid email format          | Rejected immediately with 400                        |
| Empty fields                  | Validation fails with actionable error               |
| Queue overflow                | 503 with helpful message to retry later              |

---

## 🎯 8. Developer-Facing Conventions

| Convention             | Description                          |
|------------------------|--------------------------------------|
| HTTP Method: `POST`    | Used only for creating/send requests |
| Content-Type: `JSON`   | All payloads and responses           |
| Auth (Future)          | Bearer/JWT via `Authorization` header |
| Timeout Suggestion     | Clients should timeout after 5s      |

---

## ✅ 9. API Usability Tips

- Always validate email on client side before sending
- Retry on 503, but with exponential backoff
- Log job IDs (when available) for tracking delivery lifecycle
- Use `/health` for monitoring or readiness probes

---

## 📅 10. Planned Enhancements

- Job ID returned in response for tracking
- Swagger/OpenAPI UI for live testing
- Developer dashboard (basic UI for admin/testing)
- Template support with named placeholders

---

## 🧩 Summary

This UX guidance ensures developers can integrate the email service with confidence, quickly onboard, and debug effectively. API consumers should always receive a consistent, predictable experience regardless of load or failure.

---
