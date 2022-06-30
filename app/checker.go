package app

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Checker struct{}

func NewChecker() *Checker {
	return &Checker{}
}

func (c Checker) Cost(inputHash string) (int, error) {
	hash := strings.TrimSpace(inputHash)
	return bcrypt.Cost([]byte(hash))
}

func (c Checker) Check(inputPassword string, inputHash string) error {
	hash := strings.TrimSpace(inputHash)
	passWord := strings.TrimSpace(inputPassword)
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passWord))
}
