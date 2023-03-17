package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	password = "hunter2"
	salt     = []byte{43, 155, 187, 9, 253, 53, 211, 48, 198, 192, 238, 57, 115, 148, 175, 58}
	hash     = []byte{
		107, 67, 161, 132, 16, 132, 22, 65, 44, 146, 155, 4, 66, 250, 142, 172, 7, 229, 221, 132, 166, 61, 97, 113,
		133, 236, 166, 92, 161, 117, 131, 81,
	}
)

func TestHashPassword(t *testing.T) {
	hashed := HashPassword(password, salt)

	assert.Equal(t, hash, hashed)
}

func TestVerifyPassword(t *testing.T) {
	ok := VerifyPassword(password, salt, hash)

	assert.True(t, ok)
}

func TestIdentity(t *testing.T) {
	salt := GenerateSalt()

	hashed := HashPassword(password, salt)
	ok := VerifyPassword(password, salt, hashed)

	assert.True(t, ok)
}
