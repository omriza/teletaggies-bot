package security

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	tokenEnvVar string = "TELEGRAM_API_TOKEN"
)

// ReadTokenFromEnv ...
func ReadTokenFromEnv() (string, error) {
	log.WithFields(log.Fields{
		"envVar": tokenEnvVar,
	}).Debug("Reading environment variable")

	token := os.Getenv(tokenEnvVar)
	if token == "" {
		return "", errors.New("telegram API token is missing or empty")
	}

	return token, nil

}
