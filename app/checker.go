package app

import "golang.org/x/crypto/bcrypt"

type Checker struct{}

func NewChecker() *Checker {
	return &Checker{}
}

func (c Checker) Cost(hash string) (int, error) {
	return bcrypt.Cost([]byte(hash))
}

func (c Checker) Check(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
