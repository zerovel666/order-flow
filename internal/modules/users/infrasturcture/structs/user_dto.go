package structs

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"
)

type Role string

const (
	ClientRole Role = "CLIENT"
	SellerRole Role = "SELLER"
	AdminRole  Role = "ADMIN"
)

type User struct {
	ID         uint64
	Name       string
	Password   string
	APIKeyHash string
	Role       Role
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u User) GenApiKey() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (u User) Hash(key string) string {
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:])
}
