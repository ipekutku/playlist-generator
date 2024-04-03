package auth

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type Auth interface {
	GetClient() spotify.Client
}

type SpotifyAuth struct {
	ClientID     string
	ClientSecret string
}

var _ Auth = (*SpotifyAuth)(nil) // Ensure that SpotifyAuth implements Auth

func (s *SpotifyAuth) GetClient() spotify.Client {
	config := &clientcredentials.Config{
		ClientID:     s.ClientID,
		ClientSecret: s.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	client := spotify.Authenticator{}.NewClient(token)
	return client
}
