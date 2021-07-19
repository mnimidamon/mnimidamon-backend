package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/authorization"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type userCreateImpl struct {
	JAuth restapi.JwtAuthentication
	URCase usecase.UserRegistrationInterface
}

func (impl userCreateImpl) Handle(p authorization.RegisterUserParams) middleware.Responder {
	username, password := *p.Body.Username, p.Body.Password

	user, err := impl.URCase.RegisterUser(payload.UserCredentialsPayload{
		Username: username,
		Password: password.String(),
	})

	if IsInternalError(err) {
		return authorization.NewRegisterUserInternalServerError().
			WithPayload(ErrInternalServer)
	}

	if err != nil {
		return authorization.NewRegisterUserBadRequest().
			WithPayload(ToRestError(err))
	}

	token, err := impl.JAuth.GenerateUserToken(user.ID)
	if err != nil {
		return authorization.NewRegisterUserInternalServerError().
			WithPayload(ErrSigningToken)
	}

	rp := &modelapi.RegisterResponse{
		APIKey: token,
		User:   MapToUser(user),
	}

	return authorization.NewRegisterUserOK().
		WithPayload(rp)
}

func NewUserRegistrationHandler(ur usecase.UserRegistrationInterface, ja restapi.JwtAuthentication) authorization.RegisterUserHandler {
	return &userCreateImpl{
		JAuth: ja,
		URCase: ur,
	}
}
