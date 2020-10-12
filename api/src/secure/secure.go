package secure

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash recebe uma string e coloca um hash nela
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

//VerifyPass verifica as senhas
func VerifyPass(passHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}
