package utils

import "encoding/base64"

func Encode(plain string) string {
	source := base64.StdEncoding.EncodeToString([]byte(plain))
	encoded := Base58Encode([]byte(source))
	return string(encoded)
}
