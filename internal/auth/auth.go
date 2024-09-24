package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	authKey := headers.Get("Authorization")
	if authKey == "" {
		return "", errors.New("authorization header is missing")
	}

	vals := strings.Split(authKey, "")
	if len(vals) != 2 {
		return "", errors.New("authorization header is invalid")
	}
	if vals[0] != "Bearer" {
		return "", errors.New("authorization header is invalid")
	}
	return vals[1], nil
}
