package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/imabg/responehq/pkg/logger"
	"golang.org/x/crypto/argon2"
)

func encryptPassword(password string) string {
	ctx := context.Background()
	salt := []byte("super-secure-password")
	if _, err := rand.Read(salt); err != nil {
		logger.Error(ctx, "EncryptPassword", err)
		return ""
	}
	hash := argon2.IDKey([]byte(password), salt, 10, 32*1024, 2, 32)
	return hex.EncodeToString(hash)
}

//func VerifyPassword(password, hash string) bool {
//
//}
