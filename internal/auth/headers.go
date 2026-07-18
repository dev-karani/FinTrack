package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	// confirm authorization header
	authstring := headers.Get("Authorization")
	if authstring == "" {
		return "", errors.New("missing authorization header")
	}

	// confirm beraer token prefix
	var bearerPrefix = "Bearer "

	if !strings.HasPrefix(authstring, bearerPrefix) {
		return "", errors.New("authorization header must start with Bearer")
	}

	// get bearer Token
	token := strings.TrimPrefix(authstring, bearerPrefix)
	if token == "" {
		return "", errors.New("missing bearer token")
	}
	token = strings.TrimSpace(token)
	return token, nil
}
