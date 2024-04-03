package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zmb3/spotify"
)

func TestMockAuth_GetClient(t *testing.T) {
	mockAuth := new(MockAuth)

	mockSpotifyClient := &spotify.Client{}

	mockAuth.On("GetClient").Return(*mockSpotifyClient)

	client := mockAuth.GetClient()

	assert.NotNil(t, client)
	mockAuth.AssertExpectations(t)
}
