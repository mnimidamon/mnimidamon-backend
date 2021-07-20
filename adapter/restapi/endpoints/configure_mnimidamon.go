// This file is safe to edit. Once it exists it will not be overwritten

package endpoints

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/handlers"
	"mnimidamonbackend/domain/repository/sqliterepo"
	"mnimidamonbackend/domain/usecase/listcomputer"
	"mnimidamonbackend/domain/usecase/listgroup"
	"mnimidamonbackend/domain/usecase/listgroupcomputer"
	"mnimidamonbackend/domain/usecase/listgroupmember"
	"mnimidamonbackend/domain/usecase/listuser"
	"mnimidamonbackend/domain/usecase/managegroup"
	"mnimidamonbackend/domain/usecase/userregistration"
	"net/http"

	"mnimidamonbackend/adapter/restapi/endpoints/operations"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/authorization"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/current_user"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/invite"
)

//go:generate swagger generate server --target ..\..\restapi --name Mnimidamon --spec ..\..\..\public\spec\swagger.yaml --model-package modelapi --server-package endpoints --principal interface{}

func configureFlags(api *operations.MnimidamonAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MnimidamonAPI) http.Handler {
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

	// Setting up the database.
	db, err := sqliterepo.Initialize("../databasefiles/mnimidamon.db", &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}, false)

	if err != nil {
		panic(err)
	}

	// Setting up the repositories.
	ur := sqliterepo.NewUserRepository(db)
	cr := sqliterepo.NewComputerRepository(db)
	gr := sqliterepo.NewGroupRepository(db)
	gcr := sqliterepo.NewGroupComputerRepository(db)

	// Use case setup.
	uruc := userregistration.NewUseCase(ur)
	luuc := listuser.NewUseCase(ur)
	mguc := managegroup.NewUseCase(ur, gr)
	lguc := listgroup.NewUseCase(gr)
	lcuc := listcomputer.NewUseCase(cr)
	lgcuc := listgroupcomputer.NewUseCase(gcr)
	lgmuc := listgroupmember.NewUseCase(gr)

	// Setting up the authorization.
	ja := restapi.NewJwtAuthentication("SuperSecretKey", luuc, lguc, lcuc, lgcuc, lgmuc)

	// Applies when the "X-AUTH-KEY" header is set
	api.AuthKeyAuth = ja.UserKeyMiddleware()
	// Applies when the "X-COMP-KEY" header is set
	api.CompKeyAuth = ja.CompKeyMiddleware()

	/*
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
	*/

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

	api.GroupCreateGroupHandler = handlers.NewCreateGroupHandler(mguc, ja)
	api.CurrentUserGetCurrentUserHandler = handlers.NewGetCurrentUserHandler(ja)

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

	api.CurrentUserGetCurrentUserGroupsHandler = handlers.NewGetCurrentUserGroupsHandler(lguc, ja)

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

	api.GroupGetGroupHandler = handlers.NewGetGroupHandler(ja)

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

	api.UserGetUserHandler = handlers.NewGetUserHandler(luuc)

	api.UserGetUsersHandler = handlers.NewGetUsersHandler(luuc)

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

	api.AuthorizationLoginUserHandler = handlers.NewLoginUserHandler(uruc, ja)

	if api.AuthorizationRegisterComputerHandler == nil {
		api.AuthorizationRegisterComputerHandler = authorization.RegisterComputerHandlerFunc(func(params authorization.RegisterComputerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation authorization.RegisterComputer has not yet been implemented")
		})
	}

	api.AuthorizationRegisterUserHandler = handlers.NewUserRegistrationHandler(uruc, ja)

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

	api.ServerShutdown = func() {
		sqlConn, _ := db.DB()
		_ = sqlConn.Close()
	}

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
