package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type Config struct {
	SpotifyLoginConfig oauth2.Config
}

func SpotifyConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		SpotifyLoginConfig: oauth2.Config{
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.spotify.com/authorize",
				TokenURL: "https://accounts.spotify.com/api/token",
			},
			RedirectURL: "http://localhost:8080/callback",
			Scopes:      []string{spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail},
		},
	}
}
