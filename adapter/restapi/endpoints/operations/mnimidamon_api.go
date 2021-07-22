// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"mnimidamonbackend/adapter/restapi/endpoints/operations/authorization"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/current_user"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/invite"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/user"
)

// NewMnimidamonAPI creates a new Mnimidamon instance
func NewMnimidamonAPI(spec *loads.Document) *MnimidamonAPI {
	return &MnimidamonAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		InviteAcceptCurrentUserInviteHandler: invite.AcceptCurrentUserInviteHandlerFunc(func(params invite.AcceptCurrentUserInviteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invite.AcceptCurrentUserInvite has not yet been implemented")
		}),
		GroupCreateGroupHandler: group.CreateGroupHandlerFunc(func(params group.CreateGroupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.CreateGroup has not yet been implemented")
		}),
		InviteDeclineCurrentUserInviteHandler: invite.DeclineCurrentUserInviteHandlerFunc(func(params invite.DeclineCurrentUserInviteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invite.DeclineCurrentUserInvite has not yet been implemented")
		}),
		CurrentUserDeleteCurrentUserHandler: current_user.DeleteCurrentUserHandlerFunc(func(params current_user.DeleteCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.DeleteCurrentUser has not yet been implemented")
		}),
		BackupDownloadBackupHandler: backup.DownloadBackupHandlerFunc(func(params backup.DownloadBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.DownloadBackup has not yet been implemented")
		}),
		ComputerGetBackupLocationsHandler: computer.GetBackupLocationsHandlerFunc(func(params computer.GetBackupLocationsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetBackupLocations has not yet been implemented")
		}),
		ComputerGetCurrentComputerHandler: computer.GetCurrentComputerHandlerFunc(func(params computer.GetCurrentComputerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentComputer has not yet been implemented")
		}),
		CurrentUserGetCurrentUserHandler: current_user.GetCurrentUserHandlerFunc(func(params current_user.GetCurrentUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.GetCurrentUser has not yet been implemented")
		}),
		ComputerGetCurrentUserComputerHandler: computer.GetCurrentUserComputerHandlerFunc(func(params computer.GetCurrentUserComputerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentUserComputer has not yet been implemented")
		}),
		ComputerGetCurrentUserComputersHandler: computer.GetCurrentUserComputersHandlerFunc(func(params computer.GetCurrentUserComputersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentUserComputers has not yet been implemented")
		}),
		ComputerGetCurrentUserGroupComputersHandler: computer.GetCurrentUserGroupComputersHandlerFunc(func(params computer.GetCurrentUserGroupComputersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetCurrentUserGroupComputers has not yet been implemented")
		}),
		CurrentUserGetCurrentUserGroupsHandler: current_user.GetCurrentUserGroupsHandlerFunc(func(params current_user.GetCurrentUserGroupsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.GetCurrentUserGroups has not yet been implemented")
		}),
		InviteGetCurrentUserInviteHandler: invite.GetCurrentUserInviteHandlerFunc(func(params invite.GetCurrentUserInviteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invite.GetCurrentUserInvite has not yet been implemented")
		}),
		CurrentUserGetCurrentUserInvitesHandler: current_user.GetCurrentUserInvitesHandlerFunc(func(params current_user.GetCurrentUserInvitesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation current_user.GetCurrentUserInvites has not yet been implemented")
		}),
		GroupGetGroupHandler: group.GetGroupHandlerFunc(func(params group.GetGroupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.GetGroup has not yet been implemented")
		}),
		BackupGetGroupBackupHandler: backup.GetGroupBackupHandlerFunc(func(params backup.GetGroupBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.GetGroupBackup has not yet been implemented")
		}),
		BackupGetGroupBackupsHandler: backup.GetGroupBackupsHandlerFunc(func(params backup.GetGroupBackupsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.GetGroupBackups has not yet been implemented")
		}),
		GroupGetGroupInvitesHandler: group.GetGroupInvitesHandlerFunc(func(params group.GetGroupInvitesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.GetGroupInvites has not yet been implemented")
		}),
		GroupGetGroupMembersHandler: group.GetGroupMembersHandlerFunc(func(params group.GetGroupMembersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.GetGroupMembers has not yet been implemented")
		}),
		UserGetUserHandler: user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		}),
		UserGetUsersHandler: user.GetUsersHandlerFunc(func(params user.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUsers has not yet been implemented")
		}),
		BackupInitializeGroupBackupHandler: backup.InitializeGroupBackupHandlerFunc(func(params backup.InitializeGroupBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.InitializeGroupBackup has not yet been implemented")
		}),
		BackupInitializeGroupBackupDeletionHandler: backup.InitializeGroupBackupDeletionHandlerFunc(func(params backup.InitializeGroupBackupDeletionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.InitializeGroupBackupDeletion has not yet been implemented")
		}),
		GroupInviteUserToGroupHandler: group.InviteUserToGroupHandlerFunc(func(params group.InviteUserToGroupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation group.InviteUserToGroup has not yet been implemented")
		}),
		AuthorizationLoginUserHandler: authorization.LoginUserHandlerFunc(func(params authorization.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization.LoginUser has not yet been implemented")
		}),
		AuthorizationRegisterComputerHandler: authorization.RegisterComputerHandlerFunc(func(params authorization.RegisterComputerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation authorization.RegisterComputer has not yet been implemented")
		}),
		AuthorizationRegisterUserHandler: authorization.RegisterUserHandlerFunc(func(params authorization.RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization.RegisterUser has not yet been implemented")
		}),
		BackupRequestBackupUploadHandler: backup.RequestBackupUploadHandlerFunc(func(params backup.RequestBackupUploadParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.RequestBackupUpload has not yet been implemented")
		}),
		BackupUploadBackupHandler: backup.UploadBackupHandlerFunc(func(params backup.UploadBackupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation backup.UploadBackup has not yet been implemented")
		}),

		// Applies when the "X-AUTH-KEY" header is set
		AuthKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (auth_key) X-AUTH-KEY from header param [X-AUTH-KEY] has not yet been implemented")
		},
		// Applies when the "X-COMP-KEY" header is set
		CompKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (comp_key) X-COMP-KEY from header param [X-COMP-KEY] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*MnimidamonAPI This is OpenAPI specification for Mnimidamon backend, cross platform application for file backups. */
type MnimidamonAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AuthKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-AUTH-KEY provided in the header
	AuthKeyAuth func(string) (interface{}, error)

	// CompKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-COMP-KEY provided in the header
	CompKeyAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// InviteAcceptCurrentUserInviteHandler sets the operation handler for the accept current user invite operation
	InviteAcceptCurrentUserInviteHandler invite.AcceptCurrentUserInviteHandler
	// GroupCreateGroupHandler sets the operation handler for the create group operation
	GroupCreateGroupHandler group.CreateGroupHandler
	// InviteDeclineCurrentUserInviteHandler sets the operation handler for the decline current user invite operation
	InviteDeclineCurrentUserInviteHandler invite.DeclineCurrentUserInviteHandler
	// CurrentUserDeleteCurrentUserHandler sets the operation handler for the delete current user operation
	CurrentUserDeleteCurrentUserHandler current_user.DeleteCurrentUserHandler
	// BackupDownloadBackupHandler sets the operation handler for the download backup operation
	BackupDownloadBackupHandler backup.DownloadBackupHandler
	// ComputerGetBackupLocationsHandler sets the operation handler for the get backup locations operation
	ComputerGetBackupLocationsHandler computer.GetBackupLocationsHandler
	// ComputerGetCurrentComputerHandler sets the operation handler for the get current computer operation
	ComputerGetCurrentComputerHandler computer.GetCurrentComputerHandler
	// CurrentUserGetCurrentUserHandler sets the operation handler for the get current user operation
	CurrentUserGetCurrentUserHandler current_user.GetCurrentUserHandler
	// ComputerGetCurrentUserComputerHandler sets the operation handler for the get current user computer operation
	ComputerGetCurrentUserComputerHandler computer.GetCurrentUserComputerHandler
	// ComputerGetCurrentUserComputersHandler sets the operation handler for the get current user computers operation
	ComputerGetCurrentUserComputersHandler computer.GetCurrentUserComputersHandler
	// ComputerGetCurrentUserGroupComputersHandler sets the operation handler for the get current user group computers operation
	ComputerGetCurrentUserGroupComputersHandler computer.GetCurrentUserGroupComputersHandler
	// CurrentUserGetCurrentUserGroupsHandler sets the operation handler for the get current user groups operation
	CurrentUserGetCurrentUserGroupsHandler current_user.GetCurrentUserGroupsHandler
	// InviteGetCurrentUserInviteHandler sets the operation handler for the get current user invite operation
	InviteGetCurrentUserInviteHandler invite.GetCurrentUserInviteHandler
	// CurrentUserGetCurrentUserInvitesHandler sets the operation handler for the get current user invites operation
	CurrentUserGetCurrentUserInvitesHandler current_user.GetCurrentUserInvitesHandler
	// GroupGetGroupHandler sets the operation handler for the get group operation
	GroupGetGroupHandler group.GetGroupHandler
	// BackupGetGroupBackupHandler sets the operation handler for the get group backup operation
	BackupGetGroupBackupHandler backup.GetGroupBackupHandler
	// BackupGetGroupBackupsHandler sets the operation handler for the get group backups operation
	BackupGetGroupBackupsHandler backup.GetGroupBackupsHandler
	// GroupGetGroupInvitesHandler sets the operation handler for the get group invites operation
	GroupGetGroupInvitesHandler group.GetGroupInvitesHandler
	// GroupGetGroupMembersHandler sets the operation handler for the get group members operation
	GroupGetGroupMembersHandler group.GetGroupMembersHandler
	// UserGetUserHandler sets the operation handler for the get user operation
	UserGetUserHandler user.GetUserHandler
	// UserGetUsersHandler sets the operation handler for the get users operation
	UserGetUsersHandler user.GetUsersHandler
	// BackupInitializeGroupBackupHandler sets the operation handler for the initialize group backup operation
	BackupInitializeGroupBackupHandler backup.InitializeGroupBackupHandler
	// BackupInitializeGroupBackupDeletionHandler sets the operation handler for the initialize group backup deletion operation
	BackupInitializeGroupBackupDeletionHandler backup.InitializeGroupBackupDeletionHandler
	// GroupInviteUserToGroupHandler sets the operation handler for the invite user to group operation
	GroupInviteUserToGroupHandler group.InviteUserToGroupHandler
	// AuthorizationLoginUserHandler sets the operation handler for the login user operation
	AuthorizationLoginUserHandler authorization.LoginUserHandler
	// AuthorizationRegisterComputerHandler sets the operation handler for the register computer operation
	AuthorizationRegisterComputerHandler authorization.RegisterComputerHandler
	// AuthorizationRegisterUserHandler sets the operation handler for the register user operation
	AuthorizationRegisterUserHandler authorization.RegisterUserHandler
	// BackupRequestBackupUploadHandler sets the operation handler for the request backup upload operation
	BackupRequestBackupUploadHandler backup.RequestBackupUploadHandler
	// BackupUploadBackupHandler sets the operation handler for the upload backup operation
	BackupUploadBackupHandler backup.UploadBackupHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MnimidamonAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MnimidamonAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MnimidamonAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MnimidamonAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MnimidamonAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MnimidamonAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MnimidamonAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MnimidamonAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MnimidamonAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MnimidamonAPI
func (o *MnimidamonAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AuthKeyAuth == nil {
		unregistered = append(unregistered, "XAUTHKEYAuth")
	}
	if o.CompKeyAuth == nil {
		unregistered = append(unregistered, "XCOMPKEYAuth")
	}

	if o.InviteAcceptCurrentUserInviteHandler == nil {
		unregistered = append(unregistered, "invite.AcceptCurrentUserInviteHandler")
	}
	if o.GroupCreateGroupHandler == nil {
		unregistered = append(unregistered, "group.CreateGroupHandler")
	}
	if o.InviteDeclineCurrentUserInviteHandler == nil {
		unregistered = append(unregistered, "invite.DeclineCurrentUserInviteHandler")
	}
	if o.CurrentUserDeleteCurrentUserHandler == nil {
		unregistered = append(unregistered, "current_user.DeleteCurrentUserHandler")
	}
	if o.BackupDownloadBackupHandler == nil {
		unregistered = append(unregistered, "backup.DownloadBackupHandler")
	}
	if o.ComputerGetBackupLocationsHandler == nil {
		unregistered = append(unregistered, "computer.GetBackupLocationsHandler")
	}
	if o.ComputerGetCurrentComputerHandler == nil {
		unregistered = append(unregistered, "computer.GetCurrentComputerHandler")
	}
	if o.CurrentUserGetCurrentUserHandler == nil {
		unregistered = append(unregistered, "current_user.GetCurrentUserHandler")
	}
	if o.ComputerGetCurrentUserComputerHandler == nil {
		unregistered = append(unregistered, "computer.GetCurrentUserComputerHandler")
	}
	if o.ComputerGetCurrentUserComputersHandler == nil {
		unregistered = append(unregistered, "computer.GetCurrentUserComputersHandler")
	}
	if o.ComputerGetCurrentUserGroupComputersHandler == nil {
		unregistered = append(unregistered, "computer.GetCurrentUserGroupComputersHandler")
	}
	if o.CurrentUserGetCurrentUserGroupsHandler == nil {
		unregistered = append(unregistered, "current_user.GetCurrentUserGroupsHandler")
	}
	if o.InviteGetCurrentUserInviteHandler == nil {
		unregistered = append(unregistered, "invite.GetCurrentUserInviteHandler")
	}
	if o.CurrentUserGetCurrentUserInvitesHandler == nil {
		unregistered = append(unregistered, "current_user.GetCurrentUserInvitesHandler")
	}
	if o.GroupGetGroupHandler == nil {
		unregistered = append(unregistered, "group.GetGroupHandler")
	}
	if o.BackupGetGroupBackupHandler == nil {
		unregistered = append(unregistered, "backup.GetGroupBackupHandler")
	}
	if o.BackupGetGroupBackupsHandler == nil {
		unregistered = append(unregistered, "backup.GetGroupBackupsHandler")
	}
	if o.GroupGetGroupInvitesHandler == nil {
		unregistered = append(unregistered, "group.GetGroupInvitesHandler")
	}
	if o.GroupGetGroupMembersHandler == nil {
		unregistered = append(unregistered, "group.GetGroupMembersHandler")
	}
	if o.UserGetUserHandler == nil {
		unregistered = append(unregistered, "user.GetUserHandler")
	}
	if o.UserGetUsersHandler == nil {
		unregistered = append(unregistered, "user.GetUsersHandler")
	}
	if o.BackupInitializeGroupBackupHandler == nil {
		unregistered = append(unregistered, "backup.InitializeGroupBackupHandler")
	}
	if o.BackupInitializeGroupBackupDeletionHandler == nil {
		unregistered = append(unregistered, "backup.InitializeGroupBackupDeletionHandler")
	}
	if o.GroupInviteUserToGroupHandler == nil {
		unregistered = append(unregistered, "group.InviteUserToGroupHandler")
	}
	if o.AuthorizationLoginUserHandler == nil {
		unregistered = append(unregistered, "authorization.LoginUserHandler")
	}
	if o.AuthorizationRegisterComputerHandler == nil {
		unregistered = append(unregistered, "authorization.RegisterComputerHandler")
	}
	if o.AuthorizationRegisterUserHandler == nil {
		unregistered = append(unregistered, "authorization.RegisterUserHandler")
	}
	if o.BackupRequestBackupUploadHandler == nil {
		unregistered = append(unregistered, "backup.RequestBackupUploadHandler")
	}
	if o.BackupUploadBackupHandler == nil {
		unregistered = append(unregistered, "backup.UploadBackupHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MnimidamonAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MnimidamonAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "auth_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.AuthKeyAuth)

		case "comp_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.CompKeyAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *MnimidamonAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MnimidamonAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MnimidamonAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MnimidamonAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the mnimidamon API
func (o *MnimidamonAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MnimidamonAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/current/invites/{group_id}/accept"] = invite.NewAcceptCurrentUserInvite(o.context, o.InviteAcceptCurrentUserInviteHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/current/groups"] = group.NewCreateGroup(o.context, o.GroupCreateGroupHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/current/invites/{group_id}"] = invite.NewDeclineCurrentUserInvite(o.context, o.InviteDeclineCurrentUserInviteHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/current"] = current_user.NewDeleteCurrentUser(o.context, o.CurrentUserDeleteCurrentUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers/current/groups/{group_id}/backups/{backup_id}/download"] = backup.NewDownloadBackup(o.context, o.BackupDownloadBackupHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers/current/groups/{group_id}/backups/{backup_id}/computers"] = computer.NewGetBackupLocations(o.context, o.ComputerGetBackupLocationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers/current"] = computer.NewGetCurrentComputer(o.context, o.ComputerGetCurrentComputerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current"] = current_user.NewGetCurrentUser(o.context, o.CurrentUserGetCurrentUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers/{computer_id}"] = computer.NewGetCurrentUserComputer(o.context, o.ComputerGetCurrentUserComputerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers"] = computer.NewGetCurrentUserComputers(o.context, o.ComputerGetCurrentUserComputersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/groups/{group_id}/computers"] = computer.NewGetCurrentUserGroupComputers(o.context, o.ComputerGetCurrentUserGroupComputersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/groups"] = current_user.NewGetCurrentUserGroups(o.context, o.CurrentUserGetCurrentUserGroupsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/invites/{group_id}"] = invite.NewGetCurrentUserInvite(o.context, o.InviteGetCurrentUserInviteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/invites"] = current_user.NewGetCurrentUserInvites(o.context, o.CurrentUserGetCurrentUserInvitesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/groups/{group_id}"] = group.NewGetGroup(o.context, o.GroupGetGroupHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers/current/groups/{group_id}/backups/{backup_id}"] = backup.NewGetGroupBackup(o.context, o.BackupGetGroupBackupHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/computers/current/groups/{group_id}/backups"] = backup.NewGetGroupBackups(o.context, o.BackupGetGroupBackupsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/groups/{group_id}/invites"] = group.NewGetGroupInvites(o.context, o.GroupGetGroupInvitesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/current/groups/{group_id}/members"] = group.NewGetGroupMembers(o.context, o.GroupGetGroupMembersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{user_id}"] = user.NewGetUser(o.context, o.UserGetUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = user.NewGetUsers(o.context, o.UserGetUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/current/computers/current/groups/{group_id}/backups"] = backup.NewInitializeGroupBackup(o.context, o.BackupInitializeGroupBackupHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/current/computers/current/groups/{group_id}/backups/{backup_id}"] = backup.NewInitializeGroupBackupDeletion(o.context, o.BackupInitializeGroupBackupDeletionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/current/groups/{group_id}/invites"] = group.NewInviteUserToGroup(o.context, o.GroupInviteUserToGroupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/login"] = authorization.NewLoginUser(o.context, o.AuthorizationLoginUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/current/computers"] = authorization.NewRegisterComputer(o.context, o.AuthorizationRegisterComputerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = authorization.NewRegisterUser(o.context, o.AuthorizationRegisterUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/current/computers/current/groups/{group_id}/backups/{backup_id}"] = backup.NewRequestBackupUpload(o.context, o.BackupRequestBackupUploadHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/current/computers/current/groups/{group_id}/backups/{backup_id}/upload"] = backup.NewUploadBackup(o.context, o.BackupUploadBackupHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MnimidamonAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MnimidamonAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MnimidamonAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MnimidamonAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MnimidamonAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
