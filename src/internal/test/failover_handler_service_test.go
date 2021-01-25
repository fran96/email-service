package test

import (
	"errors"
	"testing"

	"github.com/fran96/email-service/src/internal"

	"github.com/stretchr/testify/assert"
)

func TestSendMailWithBothFailing(t *testing.T) {
	mockService := &MockMailService{}
	mockService.On("SendMail", "", "subject", "message").Return("", errors.New("ERROR"))

	mockService2 := &MockMailService{}
	mockService2.On("SendMail", "", "subject", "message").Return("", errors.New("ERROR"))

	failOverService := internal.NewFailOverHandlerService(mockService, mockService2)

	_, err := failOverService.SendMail("", "subject", "message")

	assert.Error(t, err)
	assert.EqualError(t, err, "Mail service is down")
	mockService.AssertExpectations(t)
	mockService2.AssertExpectations(t)
}

func TestSendMailWithService1Failing(t *testing.T) {
	mockService := &MockMailService{}
	mockService.On("SendMail", "", "subject", "message").Return("", errors.New("ERROR"))

	mockService2 := &MockMailService{}
	mockService2.On("SendMail", "", "subject", "message").Return("sent", nil)

	failOverService := internal.NewFailOverHandlerService(mockService, mockService2)

	resBody, err := failOverService.SendMail("", "subject", "message")

	assert.NoError(t, err)
	assert.Equal(t, "sent", resBody)
	mockService.AssertExpectations(t)
	mockService2.AssertExpectations(t)
}

func TestSendMailWithService2Failing(t *testing.T) {
	mockService := &MockMailService{}
	mockService.On("SendMail", "", "subject", "message").Return("sent", nil)

	mockService2 := &MockMailService{}
	mockService2.On("SendMail", "", "subject", "message").Return("", errors.New("ERROR"))

	failOverService := internal.NewFailOverHandlerService(mockService, mockService2)

	resBody, err := failOverService.SendMail("", "subject", "message")

	assert.NoError(t, err)
	assert.Equal(t, "sent", resBody)
}
