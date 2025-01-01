package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var (
	password = "password"
)

func TestHashPassword(t *testing.T) {
	_, err := HashPassword(password)

	if err != nil {
		t.Fatal(err)
	}

}

func TestCheckPasswordHash(t *testing.T) {
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatal(err)
	}

	check := CheckPasswordHash(password, hash)

	if !check {
		t.Fatal(bcrypt.ErrMismatchedHashAndPassword)
	}
}
