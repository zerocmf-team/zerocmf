package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

func TestJwt(t *testing.T) {
	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJrZXkiOiIxIiwiZXhwIjoxNjgxMzg3Mzk4fQ.nPRXl0nwqEqiIjV4cAojYlfxjhBmY-VLFBGwE4S90V8"

	type MyCustomClaims struct {
		UserId string `json:"key"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		fmt.Printf("%v %v", claims.UserId, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}
}
