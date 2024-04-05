package auth

import (
	"context"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	log "github.com/ipekutku/playlist-generator/pkg/log"
	"golang.org/x/oauth2"
)

// Interface for Spotify Client operations, mainly for testing purposes
type SpotifyService interface {
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	Client(ctx context.Context, token *oauth2.Token) *http.Client
}

// Implementation of SpotifyService
type SpotifyServiceImpl struct {
	oauthConfig *oauth2.Config
}

// Constructor for SpotifyServiceImpl
func NewSpotifyService(oauthConfig *oauth2.Config) *SpotifyServiceImpl {
	return &SpotifyServiceImpl{
		oauthConfig: oauthConfig,
	}
}

// Implementation of Exchange method in SpotifyService
func (s *SpotifyServiceImpl) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return s.oauthConfig.Exchange(ctx, code)
}

// Implementation of Client method in SpotifyService
func (s *SpotifyServiceImpl) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return s.oauthConfig.Client(ctx, token)
}

// Asserting SpotifyServiceImpl implements SpotifyService
var _ SpotifyService = (*SpotifyServiceImpl)(nil)

// Struct for handling Spotify authorization and authentication
type AuthHandler struct {
	Spotify     SpotifyService
	Logger      log.Logger
	oauthConfig *oauth2.Config
}

// Constructor for AuthHandler
func NewAuthHandler(spotify SpotifyService, logger log.Logger, config *oauth2.Config) *AuthHandler {
	return &AuthHandler{
		Spotify:     spotify,
		Logger:      logger,
		oauthConfig: config,
	}
}

// Handler for Spotify login
// Redirects to Spotify login page
// GET /login
// Returns 303 See Other
func (h *AuthHandler) SpotifyLogin(c *fiber.Ctx) error {
	url := h.oauthConfig.AuthCodeURL("randomstate")
	h.Logger.Info("Redirecting to Spotify login", "url", url)
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return nil
}

// Handler for Spotify callback
// Handles the callback from Spotify after login
// GET /callback
func (h *AuthHandler) SpotifyCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		h.Logger.Error("Invalid state", "state", state)
		return c.SendString("Invalid state")
	}

	code := c.Query("code")
	token, err := h.Spotify.Exchange(context.Background(), code)
	if err != nil {
		h.Logger.Error("Error getting token", "error", err)
		return c.SendString("Error getting token")
	}

	resp, err := h.Spotify.Client(context.Background(), token).Get("https://api.spotify.com/v1/me")
	if err != nil {
		h.Logger.Error("Error getting user info", "error", err)
		return c.SendString("Error getting user info")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		h.Logger.Error("Error reading user info", "error", err)
		return c.SendString("Error reading user info")
	}

	h.Logger.Info("User data retrieved successfully")
	return c.SendString(string(userData))
}
