// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCurrentUserGroupComputersHandlerFunc turns a function with the right signature into a get current user group computers handler
type GetCurrentUserGroupComputersHandlerFunc func(GetCurrentUserGroupComputersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCurrentUserGroupComputersHandlerFunc) Handle(params GetCurrentUserGroupComputersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCurrentUserGroupComputersHandler interface for that can handle valid get current user group computers params
type GetCurrentUserGroupComputersHandler interface {
	Handle(GetCurrentUserGroupComputersParams, interface{}) middleware.Responder
}

// NewGetCurrentUserGroupComputers creates a new http.Handler for the get current user group computers operation
func NewGetCurrentUserGroupComputers(ctx *middleware.Context, handler GetCurrentUserGroupComputersHandler) *GetCurrentUserGroupComputers {
	return &GetCurrentUserGroupComputers{Context: ctx, Handler: handler}
}

/* GetCurrentUserGroupComputers swagger:route GET /users/current/groups/{group_id}/computers computer getCurrentUserGroupComputers

Get computers of a group

*/
type GetCurrentUserGroupComputers struct {
	Context *middleware.Context
	Handler GetCurrentUserGroupComputersHandler
}

func (o *GetCurrentUserGroupComputers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCurrentUserGroupComputersParams()
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
