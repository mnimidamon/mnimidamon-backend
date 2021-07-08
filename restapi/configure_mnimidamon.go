// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	errors2 "errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/payload"
	"mnimidamonbackend/domain/repository/sqliterepo"
	"mnimidamonbackend/domain/usecase/userregistration"
	"mnimidamonbackend/models"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"mnimidamonbackend/restapi/operations"
	"mnimidamonbackend/restapi/operations/authorization"
	"mnimidamonbackend/restapi/operations/backup"
	"mnimidamonbackend/restapi/operations/computer"
	"mnimidamonbackend/restapi/operations/current_user"
	"mnimidamonbackend/restapi/operations/group"
	"mnimidamonbackend/restapi/operations/invite"
	"mnimidamonbackend/restapi/operations/user"
)

//go:generate swagger generate server --target ..\..\mnimidamon-backend --name Mnimidamon --spec ..\public\spec\swagger.yaml --principal interface{}

func configureFlags(api *operations.MnimidamonAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MnimidamonAPI) http.Handler {
	///////////////
	db, err := sqliterepo.Initialize("../databasefiles/mnimidamon.db", &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}, false)

	if err != nil {
		panic(err)
	}

	ud := sqliterepo.NewUserRepository(db)
	ur := userregistration.NewUseCase(ud)
	///////////////

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-AUTH-KEY" header is set
	if api.AuthKeyAuth == nil {
		api.AuthKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (auth_key) X-AUTH-KEY from header param [X-AUTH-KEY] has not yet been implemented")
		}
	}
	// Applies when the "X-COMP-KEY" header is set
	if api.CompKeyAuth == nil {
		api.CompKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (comp_key) X-COMP-KEY from header param [X-COMP-KEY] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// backup.UploadBackupMaxParseMemory = 32 << 20

	if api.InviteAcceptCurrentUserInviteHandler == nil {
		api.InviteAcceptCurrentUserInviteHandler = invite.AcceptCurrentUserInviteHandlerFunc(func(params invite.AcceptCurrentUserInviteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invite.AcceptCurrentUserInvite has not yet been implemented")
		})
	}
	if api.InviteDeclineCurrentUserInviteHandler == nil {
		api.InviteDeclineCurrentUserInviteHandler = invite.DeclineCurrentUserInviteHandlerFunc(func(params invite.DeclineCurrentUserInviteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invite.DeclineCurrentUserInvite has not yet been implemented")
		})
	}
	if api.CurrentUserDeleteCurrentUserHandler == nil {
		api.CurrentUserDeleteCurrentUserHandler = current_user.DeleteCurrentUserHandlerFunc(func(params current_user.DeleteCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.DeleteCurrentUser has not yet been implemented")
		})
	}
	if api.BackupDownloadBackupHandler == nil {
		api.BackupDownloadBackupHandler = backup.DownloadBackupHandlerFunc(func(params backup.DownloadBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.DownloadBackup has not yet been implemented")
		})
	}
	if api.ComputerGetBackupLocationsHandler == nil {
		api.ComputerGetBackupLocationsHandler = computer.GetBackupLocationsHandlerFunc(func(params computer.GetBackupLocationsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetBackupLocations has not yet been implemented")
		})
	}
	if api.CurrentUserGetCurrentUserHandler == nil {
		api.CurrentUserGetCurrentUserHandler = current_user.GetCurrentUserHandlerFunc(func(params current_user.GetCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.GetCurrentUser has not yet been implemented")
		})
	}
	if api.ComputerGetCurrentUserComputerHandler == nil {
		api.ComputerGetCurrentUserComputerHandler = computer.GetCurrentUserComputerHandlerFunc(func(params computer.GetCurrentUserComputerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentUserComputer has not yet been implemented")
		})
	}
	if api.ComputerGetCurrentUserComputersHandler == nil {
		api.ComputerGetCurrentUserComputersHandler = computer.GetCurrentUserComputersHandlerFunc(func(params computer.GetCurrentUserComputersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentUserComputers has not yet been implemented")
		})
	}
	if api.ComputerGetCurrentUserGroupComputersHandler == nil {
		api.ComputerGetCurrentUserGroupComputersHandler = computer.GetCurrentUserGroupComputersHandlerFunc(func(params computer.GetCurrentUserGroupComputersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentUserGroupComputers has not yet been implemented")
		})
	}
	if api.CurrentUserGetCurrentUserGroupsHandler == nil {
		api.CurrentUserGetCurrentUserGroupsHandler = current_user.GetCurrentUserGroupsHandlerFunc(func(params current_user.GetCurrentUserGroupsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.GetCurrentUserGroups has not yet been implemented")
		})
	}
	if api.InviteGetCurrentUserInviteHandler == nil {
		api.InviteGetCurrentUserInviteHandler = invite.GetCurrentUserInviteHandlerFunc(func(params invite.GetCurrentUserInviteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invite.GetCurrentUserInvite has not yet been implemented")
		})
	}
	if api.CurrentUserGetCurrentUserInvitesHandler == nil {
		api.CurrentUserGetCurrentUserInvitesHandler = current_user.GetCurrentUserInvitesHandlerFunc(func(params current_user.GetCurrentUserInvitesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.GetCurrentUserInvites has not yet been implemented")
		})
	}
	if api.GroupGetGroupHandler == nil {
		api.GroupGetGroupHandler = group.GetGroupHandlerFunc(func(params group.GetGroupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.GetGroup has not yet been implemented")
		})
	}
	if api.BackupGetGroupBackupHandler == nil {
		api.BackupGetGroupBackupHandler = backup.GetGroupBackupHandlerFunc(func(params backup.GetGroupBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.GetGroupBackup has not yet been implemented")
		})
	}
	if api.BackupGetGroupBackupsHandler == nil {
		api.BackupGetGroupBackupsHandler = backup.GetGroupBackupsHandlerFunc(func(params backup.GetGroupBackupsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.GetGroupBackups has not yet been implemented")
		})
	}
	if api.GroupGetGroupInvitesHandler == nil {
		api.GroupGetGroupInvitesHandler = group.GetGroupInvitesHandlerFunc(func(params group.GetGroupInvitesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.GetGroupInvites has not yet been implemented")
		})
	}
	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.UserGetUsersHandler == nil {
		api.UserGetUsersHandler = user.GetUsersHandlerFunc(func(params user.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUsers has not yet been implemented")
		})
	}
	if api.BackupInitializeGroupBackupHandler == nil {
		api.BackupInitializeGroupBackupHandler = backup.InitializeGroupBackupHandlerFunc(func(params backup.InitializeGroupBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.InitializeGroupBackup has not yet been implemented")
		})
	}
	if api.BackupInitializeGroupBackupDeletionHandler == nil {
		api.BackupInitializeGroupBackupDeletionHandler = backup.InitializeGroupBackupDeletionHandlerFunc(func(params backup.InitializeGroupBackupDeletionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.InitializeGroupBackupDeletion has not yet been implemented")
		})
	}
	if api.GroupInviteUserToGroupHandler == nil {
		api.GroupInviteUserToGroupHandler = group.InviteUserToGroupHandlerFunc(func(params group.InviteUserToGroupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.InviteUserToGroup has not yet been implemented")
		})
	}
	if api.AuthorizationLoginUserHandler == nil {
		api.AuthorizationLoginUserHandler = authorization.LoginUserHandlerFunc(func(params authorization.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization.LoginUser has not yet been implemented")
		})
	}
	if api.AuthorizationRegisterComputerHandler == nil {
		api.AuthorizationRegisterComputerHandler = authorization.RegisterComputerHandlerFunc(func(params authorization.RegisterComputerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation authorization.RegisterComputer has not yet been implemented")
		})
	}

	if api.AuthorizationRegisterUserHandler == nil {
		api.AuthorizationRegisterUserHandler = authorization.RegisterUserHandlerFunc(func(params authorization.RegisterUserParams) middleware.Responder {
			user, err := ur.RegisterUser(payload.UserCredentialsPayload{
				Username: *params.Body.Username,
				Password: params.Body.Password.String(),
			})

			if errors2.Is(err, domain.ErrUserWithUsernameAlreadyExists) {
				msg := "username is taken"
				errm := domain.ErrUserWithUsernameAlreadyExists.Error()
				return authorization.NewRegisterUserNotFound().WithPayload(&models.Error{
					Code:    &errm,
					Message: &msg,
				})
			}

			return authorization.NewRegisterUserOK().WithPayload(&models.RegisterResponse{
				APIKey: new(string),
				User:   &models.User{
					UserID:   int64(user.ID),
					Username: user.Username,
				},
			})
		})
	}
	if api.BackupRequestBackupUploadHandler == nil {
		api.BackupRequestBackupUploadHandler = backup.RequestBackupUploadHandlerFunc(func(params backup.RequestBackupUploadParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.RequestBackupUpload has not yet been implemented")
		})
	}
	if api.BackupUploadBackupHandler == nil {
		api.BackupUploadBackupHandler = backup.UploadBackupHandlerFunc(func(params backup.UploadBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.UploadBackup has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
