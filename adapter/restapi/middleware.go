package restapi

import (
	errors2 "errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtRequest "github.com/dgrijalva/jwt-go/request"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"net/http"
)

var (
	CompKeyHeader = "X-COMP-KEY"
	AuthKeyHeader = "X-AUTH-KEY"
)

var CompHeaderExtractorFilter = jwtRequest.PostExtractionFilter{
	Extractor: jwtRequest.HeaderExtractor{CompKeyHeader},
	Filter: func(s string) (string, error) {
		return s, nil
	}}

var AuthHeaderExtractorFilter = jwtRequest.PostExtractionFilter{
	Extractor: jwtRequest.HeaderExtractor{AuthKeyHeader},
	Filter: func(s string) (string, error) {
		return s, nil
	}}

func (ja jwtAuthenticationImpl) CompKeyMiddleware() func(token string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		var ct computerTokenClaims
		err := ja.ParseComputerToken(token, &ct)
		if err != nil {
			return nil, errors.New(401, fmt.Sprintf("computer token validation error: %v", err))
		}
		return ct, nil
	}
}

func (ja jwtAuthenticationImpl) UserKeyMiddleware() func(token string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		var ut userTokenClaims
		err := ja.ParseUserToken(token, &ut)
		if err != nil {
			return nil, errors.New(401, fmt.Sprintf("user token validation error: %v", err))
		}
		return ut, nil
	}
}

func (ja jwtAuthenticationImpl) ExtractUserFromApiKey(req *http.Request, callback func(um *model.User) middleware.Responder) middleware.Responder {
	// Extract the user
	utc := new(userTokenClaims)

	token, err := jwtRequest.ParseFromRequest(req, &AuthHeaderExtractorFilter, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.jwtSecret), nil
	}, jwtRequest.WithClaims(utc))

	if err != nil {
		return newInternalServerErrorResponder(ErrExtractingUserAuthToken)
	}

	if !token.Valid {
		return newUnauthorizedErrorResponder(ErrInvalidUserAuthToken)
	}

	user, err := ja.LUCase.FindById(utc.UserID)
	if err != nil {
		return newInternalServerErrorResponder(err)
	}

	// Return what the callback returns.
	return callback(user)
}

func (ja jwtAuthenticationImpl) ExtractComputerFromApiKey(req *http.Request, ownerID uint, callback func(cm *model.Computer) middleware.Responder) middleware.Responder {
	ctc := new(computerTokenClaims)
	token, err := jwtRequest.ParseFromRequest(req, &CompHeaderExtractorFilter, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.jwtSecret), nil
	}, jwtRequest.WithClaims(ctc))

	if err != nil {
		return newInternalServerErrorResponder(ErrExtractingComputerAuthToken)
	}

	if !token.Valid {
		return newUnauthorizedErrorResponder(ErrInvalidComputerAuthToken)
	}

	computer, err := ja.LCCase.FindById(ctc.ComputerID)
	if err != nil {
		return newInternalServerErrorResponder(err)
	}

	if ownerID != computer.OwnerID {
		return newBadRequestErrorResponder(nil)
	}

	// Return what the callback returns.
	return callback(computer)
}

func (ja jwtAuthenticationImpl) WithGroup(um *model.User, groupID uint, callback func(gm *model.Group) middleware.Responder) middleware.Responder {
	gm, err := ja.LGCase.FindById(groupID)
	if err != nil {
		if errors2.Is(err, domain.ErrNotFound) {
			return newBadRequestErrorResponder(nil)
		} else {
			return newInternalServerErrorResponder(err)
		}
	}

	isMember, err :=  ja.LGMCase.IsMemberOf(um.ID, gm.ID)

	if err != nil {
		return newInternalServerErrorResponder(err)
	}

	if !isMember {
		return newBadRequestErrorResponder(nil)
	}

	return callback(gm)
}

func (ja jwtAuthenticationImpl) WithInvite(um *model.User, groupID uint, callback func(im *model.Invite) middleware.Responder) middleware.Responder {
	im, err := ja.LICase.FindById(um.ID, groupID)

	if err != nil {
		if errors2.Is(err, domain.ErrNotFound) {
			return newBadRequestErrorResponder(nil)
		} else {
			return newInternalServerErrorResponder(err)
		}
	}

	return callback(im)
}

// TODO have to think about it
func (ja jwtAuthenticationImpl) WithGroupComputer(req *http.Request, callback func(um *model.GroupComputer) middleware.Responder) middleware.Responder {
	panic("unimplemented")
}
