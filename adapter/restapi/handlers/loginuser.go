package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/authorization"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type loginUserImpl struct {
	JAuth  authentication.JwtAuthentication
	URCase usecase.UserRegistrationInterface
}

func (impl *loginUserImpl) Handle(p authorization.LoginUserParams) middleware.Responder {
	username, password := *p.Body.Username, p.Body.Password.String()

	user, err := impl.URCase.ValidateUserCredentials(payload.UserCredentialsPayload{
		Password: password,
		Username: username,
	})

	if IsInternalError(err) {
		return authorization.NewLoginUserInternalServerError().
			WithPayload(ErrInternalServer)
	}

	if err != nil {
		return authorization.NewLoginUserUnauthorized().
			WithPayload(ToRestError(err))
	}

	token, err := impl.JAuth.GenerateUserToken(user.ID)
	if err != nil {
		return authorization.NewLoginUserInternalServerError().
			WithPayload(ErrSigningToken)
	}

	rp := &modelapi.RegisterResponse{
		APIKey: token,
		User:   MapToUser(user),
	}

	return authorization.NewLoginUserOK().
		WithPayload(rp)
}

func NewLoginUserHandler(uruc usecase.UserRegistrationInterface, ja authentication.JwtAuthentication) authorization.LoginUserHandler {
	return &loginUserImpl{
		JAuth:  ja,
		URCase: uruc,
	}
}
