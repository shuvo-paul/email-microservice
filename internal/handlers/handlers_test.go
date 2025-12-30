package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/shuvo-paul/email-microservice/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	HealthHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Got %d, Want %d", status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("Got %s, Want %s", rr.Body.String(), expected)
	}
}

func TestEmailHandler_Success(t *testing.T) {
	mocksvc := new(service.MockEmailService)

	reqBody := models.EmailRequest{
		To:      "test@example.org",
		Subject: "Test Subject",
		Body:    "Test Body",
	}
	mocksvc.On("Send", reqBody).Return(nil)
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/send", bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := NewEmailHandler(mocksvc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code, "Handler returned wrong status")
}

func TestEmailHandler_ValidationFailure(t *testing.T) {
	mocksvc := new(service.MockEmailService)
	handler := NewEmailHandler(mocksvc)

	tests := []struct {
		name string
		body models.EmailRequest
	}{
		{
			name: "missing email address",
			body: models.EmailRequest{
				To:      "",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
		},
		{
			name: "wrong email address",
			body: models.EmailRequest{
				To:      "text@example",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
		},
		{
			name: "missing subject",
			body: models.EmailRequest{
				To:      "test@example.org",
				Subject: "",
				Body:    "Test Body",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/send", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusBadRequest, rr.Code, "status code didn't match")
			mocksvc.AssertNotCalled(t, "Send", nil)
		})
	}
}

func TestSendEmail_ServiceError(t *testing.T) {
	mocksvg := new(service.MockEmailService)
	mocksvg.On("Send", mock.Anything).Return(service.ErrSendingEmail)

	reqBody := models.EmailRequest{
		To:      "test@example.org",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/send", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := NewEmailHandler(mocksvg)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "wrong status code")
	assert.Contains(t, rr.Body.String(), service.ErrSendingEmail.Error(), "wrong error message")
}
