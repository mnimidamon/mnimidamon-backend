package restapi

import (
	"fmt"
	"github.com/go-openapi/errors"
)

func (ja jwtAuthenticationImpl) CompKeyMiddleware() func(token string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		var ct computerTokenClaims
		err := ja.ParseComputerToken(token, &ct)
		if err != nil {
			return nil, errors.New(401, fmt.Sprintf("computer token validation error: %v", err))
		}
		return nil, nil
	}
}

func (ja jwtAuthenticationImpl) UserKeyMiddleware() func(token string) (interface {}, error) {
	return func(token string) (interface{}, error) {
		var ut userTokenClaims
		err := ja.ParseUserToken(token, &ut)
		if err != nil {
			return nil, errors.New(401, fmt.Sprintf("user token validation error: %v", err))
		}
		return nil, nil
	}
}

