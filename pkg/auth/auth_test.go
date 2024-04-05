package auth

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	mocks "github.com/ipekutku/playlist-generator/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

func TestSpotifyLogin(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Mock SpotifyService
	spotifyMock := new(mocks.SpotifyServiceMock)

	// Mock Logger
	loggerMock := new(mocks.LoggerMock) // Assuming you have a LoggerMock
	loggerMock.On("Info", mock.Anything, mock.Anything).Return()

	// Initialize AuthHandler with the mock
	authHandler := NewAuthHandler(spotifyMock, loggerMock, &oauth2.Config{})

	// Define the route
	app.Get("/login", authHandler.SpotifyLogin)

	// Create a request to pass to our handler
	req := httptest.NewRequest("GET", "/login", nil)

	// Create a ResponseRecorder to record the response
	resp, _ := app.Test(req, -1)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusFound, resp.StatusCode)

	// Verify that the mock was called
	spotifyMock.AssertExpectations(t)
}

func TestSpotifyCallback(t *testing.T) {
	// Setup
	app := fiber.New()
	spotifyMock := new(mocks.SpotifyServiceMock)
	loggerMock := new(mocks.LoggerMock)

	token := &oauth2.Token{AccessToken: "access_token"}
	spotifyMock.On("Exchange", mock.Anything, "valid_code").Return(token, nil)

	httpClientMock := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) *http.Response {
			// Simulate a response from the Spotify API
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"user":"data"}`)),
				Header:     make(http.Header),
			}
		}),
	}
	spotifyMock.On("Client", mock.Anything, token).Return(httpClientMock)

	loggerMock.On("Info", mock.Anything, mock.Anything).Once()
	loggerMock.On("Error", mock.Anything, mock.Anything).Maybe() // Use .Maybe() for calls that might or might not happen

	authHandler := NewAuthHandler(spotifyMock, loggerMock, &oauth2.Config{})

	// Define the route
	app.Post("/callback", authHandler.SpotifyCallback) // Ensure the method matches your actual route definition

	// Execute the request
	req := httptest.NewRequest("POST", "/callback?state=randomstate&code=valid_code", nil)
	resp, _ := app.Test(req, -1)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify that the mock's expectations were met
	spotifyMock.AssertExpectations(t)
	loggerMock.AssertExpectations(t)
}

// Helper for mocking HTTP client responses
type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
