package mocks

import (
	"github.com/stretchr/testify/mock"
)

// LoggerMock is a mock type for the Logger interface
type LoggerMock struct {
	mock.Mock
}

// Info mocks the Info method of the Logger interface
func (m *LoggerMock) Info(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

// Error mocks the Error method of the Logger interface
func (m *LoggerMock) Error(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

// Sync mocks the Sync method of the Logger interface
func (m *LoggerMock) Sync() {
	m.Called()
}
