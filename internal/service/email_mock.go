package service

import (
	"github.com/shuvo-paul/email-microservice/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) Send(email models.EmailRequest) error {
	args := m.Called(email)
	return args.Error(0)
}
