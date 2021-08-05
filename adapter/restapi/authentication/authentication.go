package authentication

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang-jwt/jwt"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
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
	ExtractComputerFromApiKey(req *http.Request, ownerID uint, callback func(cm *model.Computer) middleware.Responder) middleware.Responder
	ExtractUserFromApiKey(req *http.Request, callback func(um *model.User) middleware.Responder) middleware.Responder
	WithGroup(um *model.User, groupID uint, callback func(gm *model.Group) middleware.Responder) middleware.Responder
	WithInvite(um *model.User, groupID uint, callback func(im *model.Invite) middleware.Responder) middleware.Responder
	WithComputer(um *model.User, computerID uint, callback func(cm *model.Computer) middleware.Responder) middleware.Responder
	WithGroupComputer(cm *model.Computer, gm *model.Group, callback func(gcm *model.GroupComputer) middleware.Responder) middleware.Responder
	WithBackup(um *model.User, gm *model.Group, backupID uint, callback func(bm *model.Backup) middleware.Responder) middleware.Responder
}

type jwtAuthenticationImpl struct {
	LUCase  usecase.ListUserInterface
	LCCase  usecase.ListComputerInterface
	LGCase  usecase.ListGroupInterface
	LGCCase usecase.ListGroupComputerInterface
	LGMCase usecase.ListGroupMemberInterface
	LICase  usecase.ListInviteInterface
	LBCase  usecase.ListBackupInterface

	jwtSecret string
}

func NewJwtAuthentication(jwtSecret string, luuc usecase.ListUserInterface, lguc usecase.ListGroupInterface, lcuc usecase.ListComputerInterface, lgcuc usecase.ListGroupComputerInterface, lgmuc usecase.ListGroupMemberInterface, liuc usecase.ListInviteInterface, lbuc usecase.ListBackupInterface) JwtAuthentication {
	return &jwtAuthenticationImpl{
		jwtSecret: jwtSecret,
		LUCase:    luuc,
		LCCase:    lcuc,
		LGCase:    lguc,
		LGCCase:   lgcuc,
		LGMCase:   lgmuc,
		LICase:    liuc,
		LBCase:    lbuc,
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
	expiresAt := time.Unix(1<<63-1, 0)

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
	expiresAt := time.Unix(1<<63-1, 0)

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
