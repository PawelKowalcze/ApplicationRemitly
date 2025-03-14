package auth

import (
	"errors"
	"strings"
)

func GetSWIFTCode(urlPath string) (string, error) {
	parts := strings.Split(urlPath, "/")
	if len(parts) < 3 {
		return "", errors.New("invalid URL path")
	}
	swiftCode := parts[len(parts)-1]
	if swiftCode == "" {
		return "", errors.New("SWIFT code not found in URL path")
	}
	return swiftCode, nil
}

func GetCountryCode(urlPath string) (string, error) {
	parts := strings.Split(urlPath, "/")
	if len(parts) < 4 {
		return "", errors.New("invalid URL path")
	}
	countryCode := parts[len(parts)-1]
	if countryCode == "" {
		return "", errors.New("country code not found in URL path")
	}
	return countryCode, nil
}
