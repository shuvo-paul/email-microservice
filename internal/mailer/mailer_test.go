package mailer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockSender struct{}

func (m *mockSender) Send(to, subject, body string) error {
	return nil
}

func TestSMTPMailer_InputValidation(t *testing.T) {
	mailer := NewClient("smtp.gmail.com", 597, "user", "pass", "from@example.org")
	mailer.sender = &mockSender{}

	tests := []struct {
		name      string
		to        string
		subject   string
		body      string
		expectErr bool
		errMsg    string
	}{
		{
			name:      "valid recipient email",
			to:        "user@example.org",
			subject:   "Normal Subject",
			body:      "Normal body content",
			expectErr: false,
		},
		{
			name:      "invalid recipient email",
			to:        "not-an-email",
			subject:   "subject",
			body:      "body",
			expectErr: true,
			errMsg:    "invalid recipient address",
		},
		{
			name:      "subject with newline - header injection attack",
			to:        "user@example.com",
			subject:   "Subject\nBcc: hacker@evil.com",
			body:      "Body",
			expectErr: true,
			errMsg:    "subject cannot contain newlines",
		},
		{
			name:      "subject with CRLF - header injection attack",
			to:        "user@example.com",
			subject:   "Subject\r\nBcc: hacker@evil.com",
			body:      "Body",
			expectErr: true,
			errMsg:    "subject cannot contain newlines",
		},
		{
			name:      "extremely long subject",
			to:        "user@example.com",
			subject:   strings.Repeat("A", 300),
			body:      "Body",
			expectErr: true,
			errMsg:    "subject too long",
		},
		{
			name:      "subject with control characters",
			to:        "user@example.com",
			subject:   "Subject with \x00 null byte",
			body:      "Body",
			expectErr: false, // Should be cleaned, not error
		},
		{
			name:      "empty fields are OK",
			to:        "user@example.com",
			subject:   "",
			body:      "",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mailer.Send(tt.to, tt.subject, tt.body)
			if tt.expectErr {
				assert.ErrorContains(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
