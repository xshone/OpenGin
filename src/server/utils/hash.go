package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func HashText(text string) string {
	md5Handler := md5.New()
	md5Handler.Write([]byte(text))
	hashedBytes := md5Handler.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}

func HashPassword(password, salt string) string {
	text := password + salt
	saltParts := strings.Split(salt, "-")

	if len(saltParts) > 1 {
		text = saltParts[0] + password + saltParts[len(saltParts)-1]
	}

	hashedPassword := HashText(text)
	return hashedPassword
}
