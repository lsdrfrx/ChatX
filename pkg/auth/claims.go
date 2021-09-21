package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.StandardClaims
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
}
