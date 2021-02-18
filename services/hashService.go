package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashService Interface defining hash service
type HashService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type hashService struct {
	salt int
}

// NewHashService a function to create a new hash service
func NewHashService(salt int) HashService {
	return &hashService{
		salt,
	}
}

// HashPassword method to hash a password
func (s *hashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), s.salt)
	return string(bytes), err
}

// CheckPasswordHash method to check a password and hash
func (s *hashService) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err, hash, password)
	return err == nil
}
