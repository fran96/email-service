package test

import (
	"github.com/stretchr/testify/mock"
)
type MockMailService struct {
	mock.Mock
}

func (mg *MockMailService) SendMail(to, subject, message string) (string, error) {
	args := mg.Called(to, subject, message)
	return args.String(0), args.Error(1)
}
