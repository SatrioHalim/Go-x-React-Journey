package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Secret key
	secret := []byte("INI_ADALAH_SECRET_TOKEN")
	// Isi payload JWT
	claims := jwt.MapClaims{
		"uid" : "1234",
		"role" : "admin",
		"exp" : time.Now().Add(time.Hour).Unix(), // ambil current time ditambah satu jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims);
	
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generate JWT Token")
	fmt.Println(signedToken)
}