package auth

import jwt "github.com/dgrijalva/jwt-go"

//JwtClaims jwt
type JwtClaims struct {
	ID    int      `json:"id"`
	Scope []string `json:"scope"`
	jwt.StandardClaims
}
