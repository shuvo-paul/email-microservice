package mailer

import "github.com/stretchr/testify/mock"

type mockClient struct {
	mock.Mock
}

func (m *mockClient) Send(to, subject, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}
