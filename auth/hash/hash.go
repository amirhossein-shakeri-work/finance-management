package hash

import (
	"crypto/sha512"
	"encoding/hex"
)

func GenerateHash (str string) string {
	h := sha512.Sum384([]byte(str))
	return hex.EncodeToString(h[:])
}

func ValidateHash (hash string, str string) bool {
	return GenerateHash(str) == hash
}
