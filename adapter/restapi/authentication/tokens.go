package authentication

import (
	"errors"
)

// This will be saved inside our token for user authentication
type userTokenClaims struct {
	UserID uint   `json:"uid"`
	Issuer string `json:"iss"`
}

func (ut userTokenClaims) Valid() error {
	if ut.Issuer != "mnimidamon-server" {
		return errors.New("not issued by mnimidamon server")
	}
	return nil
}

// This will be saved inside our token for computer authentication.
type computerTokenClaims struct {
	ComputerID uint   `json:"cid"`
	Issuer     string `json:"iss"`
}

func (ct computerTokenClaims) Valid() error {
	if ct.Issuer != "mnimidamon-server" {
		return errors.New("not issued by mnimidamon server")
	}
	return nil
}
