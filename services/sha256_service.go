package services

import (
	"crypto/sha256"
	"fmt"
)

// funcao para encriptar
func SHA256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)
}