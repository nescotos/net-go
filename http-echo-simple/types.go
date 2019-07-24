package main

import "github.com/dgrijalva/jwt-go"

type jwtLoginClaims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.StandardClaims
}
