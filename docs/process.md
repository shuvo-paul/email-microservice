# 🔄 Development & Maintenance Process
## Email Notification Microservice

---

## 📌 1. Overview

This document defines the development practices, lifecycle workflows, testing strategy, and CI/CD pipeline used to develop and maintain the Email Notification Microservice.

---

## 🔁 2. Branching Strategy

We follow a simplified Git workflow:

- `main`: stable production-ready code
- `dev`: active development and integration
- `feature/*`: new features or enhancements
- `bugfix/*`: patches or hotfixes

### Git Commands

```bash
git checkout -b feature/email-templates
# work...
git push origin feature/email-templates
```

Pull Requests (PRs) must target `dev`. `main` is updated only via tested releases.

---

## 🧪 3. Testing Strategy

### Unit Testing

- Focus on job validation, worker behavior, and retry logic
- Use Go's `testing` package
- Optional: integrate with `stretchr/testify`

### Manual Testing

- Using `curl`, Postman, or scripts
- Validate inputs, outputs, and retry behaviors

### Test Coverage

- All PRs must maintain or improve test coverage
- Run tests locally before committing:
```bash
go test ./...
```

---

## 💡 4. Linting & Formatting

Use `gofmt` and `golangci-lint` to maintain code quality:

```bash
gofmt -s -w .
golangci-lint run
```

Linting runs automatically via CI.

---

## 🚦 5. CI/CD Workflow

CI tasks include:

- ✅ Format & Lint
- ✅ Run Tests
- ✅ Build Docker Image
- ✅ Publish to Registry (planned)
- ✅ Deploy to Staging/Production (planned)

### Example GitHub Actions (planned)

```yaml
on: [push]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - run: go test ./...
```

---

## 🧹 6. Code Reviews

- Review your own code for logic, performance, and security issues before merging.
- Use comments in code or commit messages to explain important decisions or tricky parts.
- For major features or bugs, consider tracking them with issues or PRs for documentation, but this is optional for solo development.

---

## 🧰 7. Issue Management

| Label        | Description                  |
|--------------|------------------------------|
| `bug`        | Something isn't working       |
| `enhancement`| New feature or improvement    |
| `question`   | Clarification needed          |
| `wontfix`    | Not planned to be addressed   |

---

## 🧑‍💻 8. Release Strategy

- No formal versioning or release process is used for this solo project.
- Features and changes are tracked informally in the documentation.

---

## 📎 9. Documentation Process

- Markdown files (`*.md`) stored in root/docs
- Diagrams: Draw.io, Mermaid.js, or PlantUML
- API reference: `openapi.yaml`

---

## 🧪 10. Maintenance Policy

- Bugs are triaged weekly
- Dependencies updated monthly
- Monitoring and alerting handled externally (Prometheus, etc.)

---

## ✅ Summary

This process ensures that the project remains stable, testable, and developer-friendly while maintaining high standards for maintainability and reliability.

---
