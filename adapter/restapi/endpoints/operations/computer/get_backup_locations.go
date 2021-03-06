// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBackupLocationsHandlerFunc turns a function with the right signature into a get backup locations handler
type GetBackupLocationsHandlerFunc func(GetBackupLocationsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBackupLocationsHandlerFunc) Handle(params GetBackupLocationsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetBackupLocationsHandler interface for that can handle valid get backup locations params
type GetBackupLocationsHandler interface {
	Handle(GetBackupLocationsParams, interface{}) middleware.Responder
}

// NewGetBackupLocations creates a new http.Handler for the get backup locations operation
func NewGetBackupLocations(ctx *middleware.Context, handler GetBackupLocationsHandler) *GetBackupLocations {
	return &GetBackupLocations{Context: ctx, Handler: handler}
}

/* GetBackupLocations swagger:route GET /users/current/computers/current/groups/{group_id}/backups/{backup_id}/computers computer getBackupLocations

Get a list on which computers the backup is stored

*/
type GetBackupLocations struct {
	Context *middleware.Context
	Handler GetBackupLocationsHandler
}

func (o *GetBackupLocations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBackupLocationsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
