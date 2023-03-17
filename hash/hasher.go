package hash

import (
	"bytes"
	"crypto/rand"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/argon2"
)

const (
	saltLength = 16
	time       = 1
	memory     = 64 * 1024
	threads    = 4
	keyLength  = 32
)

func GenerateSalt() []byte {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		log.Panic().Str("tag", "hasher").Err(err).Msg("failed to generate salt")
	}
	return salt
}

func HashPassword(pw string, salt []byte) []byte {
	return argon2.IDKey([]byte(pw), salt, time, memory, threads, keyLength)
}

func VerifyPassword(pw string, salt []byte, against []byte) bool {
	key2 := HashPassword(pw, salt)

	if len(key2) != len(against) {
		return false
	}

	return bytes.Equal(against, key2)
}
