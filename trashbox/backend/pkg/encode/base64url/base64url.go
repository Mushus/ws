package base64url

import (
	"encoding/base64"
	"strings"
)

var (
	encodeReplacer = strings.NewReplacer("=", "-", "/", "_", "+", ".")
	decodeReplacer = strings.NewReplacer("-", "=", "_", "/", ".", "+")
)

func Encode(raw string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(raw))
	return encodeReplacer.Replace(encoded)
}

func Decode(encoded string) string {
	encoded = decodeReplacer.Replace(encoded)
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(b)
}
