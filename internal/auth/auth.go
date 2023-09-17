package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {api key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("Malformed Authorization info")
	}
	return vals[1], nil
}
