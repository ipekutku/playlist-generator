package auth

import (
	"github.com/stretchr/testify/mock"
	"github.com/zmb3/spotify"
)

// A mock type for the Auth interface
type MockAuth struct {
	mock.Mock
}

// GetClient provides a mock function
// Returns a spotify.Client
func (m *MockAuth) GetClient() spotify.Client {
	args := m.Called()
	return args.Get(0).(spotify.Client)
}
