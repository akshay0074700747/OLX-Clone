package helpers

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func Hash_pass(pass string) string {
	password := []byte(pass)

	// bcrypt.DefaultCost specifies how many iterations of the key derivation function (bcrypt) should be used.
	// The bcrypt.DefaultCost is a predefined constant that usually corresponds to a reasonable and secure value (typically 10).
	//we can adjust the cost factor to make the hash more or less computationally intensive.
	//Increasing the cost factor makes it more secure but slower.

	hashedpass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic("Password cannot be hashed...")
	}
	return string(hashedpass)
}

func Generate_Uid() (string, error) {
	random := make([]byte, 8)
	
	if _, err := rand.Read(random); err != nil {
		return "", err
	}

	uid := base64.URLEncoding.EncodeToString(random)

	return uid, nil
}
