package public

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

// ハッシュ化するための関数
func TokenHash(token string) string {
	b := make([]byte, 14)

	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		fmt.Println("error:", err)
	}

	salt := base64.StdEncoding.EncodeToString(b)
	converted, _ := scrypt.Key([]byte(token), []byte(salt), 32768, 8, 1, 32)

	hash := hex.EncodeToString(converted[:])
	return hash
}
