// Package handlers
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/shuvo-paul/email-microservice/internal/service"
	"github.com/shuvo-paul/email-microservice/pkg/validator"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

type EmailHandler struct {
	srvc service.Sender
}

func NewEmailHandler(emailService service.Sender) *EmailHandler {
	return &EmailHandler{
		srvc: emailService,
	}
}

func (h *EmailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := validator.ValidateEmailRequest(req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := h.srvc.Send(req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusAccepted, map[string]string{"message": "Email queued for sending"})
}

func writeJSON(w http.ResponseWriter, status int, payload map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
