package helper

import (
	"encoding/base64"
	"encoding/json"
	"io"
)

// ParseBody populates the interface provided using the request body(io.ReadCloser) to a the struct
func ParseBody(bdy io.ReadCloser, s interface{}) error {
	return json.NewDecoder(bdy).Decode(&s)
}

// EncodeToBase64 encodes a given string to base64
func EncodeToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// IIsBase64Encoded determines where the given string is already base64 encoded or not
func IsBase64Encoded(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	return true
}
