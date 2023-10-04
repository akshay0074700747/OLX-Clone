package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)


type Payload struct {
	Username string
	Email string
	jwt.StandardClaims
}


func Generatejwt(username,email string,secret []byte) (string, error) {

	expireAt := time.Now().Add(48 * time.Hour)

	jwtclaims := &Payload{
		Username: username,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtclaims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString,nil 

}


func ValidateToken(singnedToken string,secret []byte) (map[string]string , error) {

	token, err := jwt.ParseWithClaims(
		singnedToken,
		&Payload{},
		func(token *jwt.Token) (interface{}, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil,fmt.Errorf("wrong signing method")
			}

			return secret, nil

		},
	)


	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Payload)

	cred := map[string]string{
		"username" : claims.Username,
		"email" : claims.Email,
	}

	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}

	return cred,nil
	
}
