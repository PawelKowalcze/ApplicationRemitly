package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetSWIFTCode extracts the SWIFTCode from the headers
// Example:
// Code: SWIFTCode {insert SWIFTCode here}
func GetSWIFTCode(headers http.Header) (string, error) {
	val := headers.Get("Code")
	if val == "" {
		return "", errors.New("code header not found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid code header")
	}
	if vals[0] != "SWIFTCode" {
		return "", errors.New("malformed first auth header")
	}
	return vals[1], nil
}
