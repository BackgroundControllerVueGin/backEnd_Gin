package model

import "github.com/dgrijalva/jwt-go"

type UserToken struct {
	No       string `json:"no"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
