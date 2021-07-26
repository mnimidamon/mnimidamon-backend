package restapi

import "github.com/golang-jwt/jwt"

// This will be saved inside our token for user authentication
type userTokenClaims struct {
	UserID         uint       `json:"user_id"`
	StandardClaims jwt.StandardClaims `json:"standard_claims"`
}

func (ut userTokenClaims) Valid() error {
	return ut.StandardClaims.Valid()
}

// This will be saved inside our token for computer authentication.
type computerTokenClaims struct {
	ComputerID         uint       `json:"user_id"`
	StandardClaims jwt.StandardClaims `json:"standard_claims"`
}

func (ct computerTokenClaims) Valid() error {
	return ct.StandardClaims.Valid()
}