// This file is safe to edit. Once it exists it will not be overwritten

package endpoints

import (
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/current_user"
	"mnimidamonbackend/adapter/restapi/handlers"
	"mnimidamonbackend/domain/repository/filestore"
	"mnimidamonbackend/domain/repository/sqliterepo"
	"mnimidamonbackend/domain/usecase/computerregistration"
	"mnimidamonbackend/domain/usecase/groupinvite"
	"mnimidamonbackend/domain/usecase/listbackup"
	"mnimidamonbackend/domain/usecase/listcomputer"
	"mnimidamonbackend/domain/usecase/listgroup"
	"mnimidamonbackend/domain/usecase/listgroupcomputer"
	"mnimidamonbackend/domain/usecase/listgroupmember"
	"mnimidamonbackend/domain/usecase/listinvite"
	"mnimidamonbackend/domain/usecase/listuser"
	"mnimidamonbackend/domain/usecase/managebackup"
	"mnimidamonbackend/domain/usecase/managefile"
	"mnimidamonbackend/domain/usecase/managegroup"
	"mnimidamonbackend/domain/usecase/managegroupcomputer"
	"mnimidamonbackend/domain/usecase/userregistration"
	"net/http"
	"time"
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

	// File store setup.
	fs := filestore.New("../filestore")

	// Setting up the repositories.
	ur := sqliterepo.NewUserRepository(db)
	cr := sqliterepo.NewComputerRepository(db)
	gr := sqliterepo.NewGroupRepository(db)
	gcr := sqliterepo.NewGroupComputerRepository(db)
	ir := sqliterepo.NewInviteRepository(db)
	br := sqliterepo.NewBackupRepository(db)
	cbr := sqliterepo.NewComputerBackupRepository(db)

	// Use case setup.
	uruc := userregistration.NewUseCase(ur)
	luuc := listuser.NewUseCase(ur)
	mguc := managegroup.NewUseCase(ur, gr)
	lguc := listgroup.NewUseCase(gr)
	lcuc := listcomputer.NewUseCase(cr)
	lgcuc := listgroupcomputer.NewUseCase(gcr)
	lgmuc := listgroupmember.NewUseCase(gr)
	crcuc := computerregistration.NewUseCase(cr, ur)
	giuc := groupinvite.NewUseCase(gr, ir, ur)
	liuc := listinvite.NewUseCase(ir)
	mgcuc := managegroupcomputer.NewUseCase(gcr, cr, gr, br, cbr)
	mbuc := managebackup.NewUseCase(fs, br, ur, gr, cr, gcr, cbr)
	lbuc := listbackup.NewUseCase(br)
	mfuc := managefile.NewUseCase(fs, br)

	// Setting up the authorization.
	ja := restapi.NewJwtAuthentication("SuperSecretKey", luuc, lguc, lcuc, lgcuc, lgmuc, liuc, lbuc)

	// Applies when the "X-AUTH-KEY" header is set
	api.AuthKeyAuth = ja.UserKeyMiddleware()
	// Applies when the "X-COMP-KEY" header is set
	api.CompKeyAuth = ja.CompKeyMiddleware()

	api.AuthorizationLoginUserHandler = handlers.NewLoginUserHandler(uruc, ja)
	api.AuthorizationRegisterComputerHandler = handlers.NewRegisterComputerHandler(crcuc, ja)
	api.AuthorizationRegisterUserHandler = handlers.NewUserRegistrationHandler(uruc, ja)

	api.InviteAcceptCurrentUserInviteHandler = handlers.NewAcceptInviteHandler(giuc, ja)
	api.InviteDeclineCurrentUserInviteHandler = handlers.NewDeclineCurrentUserInviteHandler(giuc, ja)

	api.GroupGetGroupMembersHandler = handlers.NewGetGroupMembersHandler(lgmuc, ja)
	api.GroupComputerJoinComputerToGroupHandler = handlers.NewJoinComputerToGroupHandler(mgcuc, ja)
	api.GroupInviteUserToGroupHandler = handlers.NewInviteUserToGroupHandler(giuc, luuc, ja)
	api.GroupCreateGroupHandler = handlers.NewCreateGroupHandler(mguc, ja)
	api.GroupGetGroupHandler = handlers.NewGetGroupHandler(ja)
	api.GroupGetGroupInvitesHandler = handlers.NewGetGroupInvitesHandler(liuc, ja)

	api.UserGetUserHandler = handlers.NewGetUserHandler(luuc)
	api.UserGetUsersHandler = handlers.NewGetUsersHandler(luuc)

	api.CurrentUserGetCurrentUserHandler = handlers.NewGetCurrentUserHandler(ja)
	api.CurrentUserGetCurrentUserInvitesHandler = handlers.NewGetCurrentUserInvitesHandler(liuc, ja)
	api.CurrentUserGetCurrentUserGroupsHandler = handlers.NewGetCurrentUserGroupsHandler(lguc, ja)

	api.ComputerGetCurrentComputerHandler = handlers.NewGetCurrentUserComputer(ja)
	api.ComputerGetCurrentUserComputerHandler = handlers.NewGetCurrentUserComputerHandler(lcuc, ja)

	api.InviteGetCurrentUserInviteHandler = handlers.NewGetCurrentUserInviteHandler(ja)

	api.BackupInitializeGroupBackupHandler = handlers.NewInitializeGroupBackupHandler(mbuc, ja)
	api.BackupInitializeGroupBackupDeletionHandler = handlers.NewGroupBackupDeletionImpl(mbuc, ja)
	api.BackupUploadBackupHandler = handlers.NewUploadBackupHandler(mfuc, ja)

	api.GroupComputerLeaveComputerFromGroupHandler = nil

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

	if api.BackupRequestBackupUploadHandler == nil {
		api.BackupRequestBackupUploadHandler = backup.RequestBackupUploadHandlerFunc(func(params backup.RequestBackupUploadParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.RequestBackupUpload has not yet been implemented")
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
	// Request logging middleware.
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		handler.ServeHTTP(rw, r) // serve the original request

		duration := time.Since(start)

		// log request details
		fmt.Printf("%v %v %v - %v\n", start.Format("2006/01/02 15:04:05"), method, uri, duration)
	}

	return http.HandlerFunc(logFn)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
