package restapi

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"net/http"
	"strconv"
	"time"
)

type JwtAuthentication interface {
	ParseComputerToken(tokenString string, claims *computerTokenClaims) error
	ParseUserToken(tokenString string, claims *userTokenClaims) error
	GenerateComputerToken(computerID uint) (*string, error)
	GenerateUserToken(userID uint) (*string, error)
	CompKeyMiddleware() func(token string) (interface{}, error)
	UserKeyMiddleware() func(token string) (interface{}, error)
	ExtractComputerFromApiKey(req *http.Request, callback func(um *model.Computer) middleware.Responder) middleware.Responder
	ExtractUserFromApiKey(req *http.Request, callback func(um *model.User) middleware.Responder) middleware.Responder
}

type jwtAuthenticationImpl struct {
	URepo  repository.UserRepository
	CRepo  repository.ComputerRepository
	GCRepo repository.GroupComputerRepository

	jwtSecret string
}

func NewJwtAuthentication(jwtSecret string, ur repository.UserRepository, cr repository.ComputerRepository, gcr repository.GroupComputerRepository) JwtAuthentication {
	return &jwtAuthenticationImpl{
		jwtSecret: jwtSecret,
		URepo:     ur,
		CRepo:     cr,
		GCRepo:    gcr,
	}
}

func (ja *jwtAuthenticationImpl) ParseComputerToken(tokenString string, claims *computerTokenClaims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.jwtSecret), nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*computerTokenClaims); ok && !token.Valid {
		return ErrInvalidComputerAuthToken
	}

	return nil
}

func (ja *jwtAuthenticationImpl) ParseUserToken(tokenString string, claims *userTokenClaims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ja.jwtSecret), nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*userTokenClaims); ok && !token.Valid {
		return ErrInvalidUserAuthToken
	}

	return nil
}

func (ja *jwtAuthenticationImpl) GenerateComputerToken(computerID uint) (*string, error) {
	// The tokens will expire in one day. Unix function converts the
	// date to the seconds passed so int64.
	expiresAt := time.Now().Add(time.Hour * 10)

	// Populate the claims.
	claims := computerTokenClaims{
		ComputerID: computerID,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.FormatInt(int64(computerID), 10),
			Issuer:    "mnimidamon-server",
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	return ja.generateSignedString(claims)
}

func (ja *jwtAuthenticationImpl) generateSignedString(claims jwt.Claims) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the JWT_KEY environment variable.
	key := []byte(ja.jwtSecret)
	signedString, err := token.SignedString(key)

	if err != nil {
		return nil, err
	}

	return &signedString, nil
}

func (ja *jwtAuthenticationImpl) GenerateUserToken(userID uint) (*string, error) {
	// The tokens will expire in one day. Unix function converts the
	// date to the seconds passed so int64.
	expiresAt := time.Now().Add(time.Hour * 10)

	// Populate the claims.
	claims := userTokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.FormatInt(int64(userID), 10),
			Issuer:    "mnimidamon-server",
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	return ja.generateSignedString(claims)
}
