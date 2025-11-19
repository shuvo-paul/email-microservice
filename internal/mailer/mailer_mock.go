package mailer

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Send(to, subject, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}

type MockSender struct {
	mock.Mock
}

func (m *MockSender) Send(to, subject, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}
