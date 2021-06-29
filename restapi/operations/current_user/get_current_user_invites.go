// Code generated by go-swagger; DO NOT EDIT.

package current_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCurrentUserInvitesHandlerFunc turns a function with the right signature into a get current user invites handler
type GetCurrentUserInvitesHandlerFunc func(GetCurrentUserInvitesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCurrentUserInvitesHandlerFunc) Handle(params GetCurrentUserInvitesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCurrentUserInvitesHandler interface for that can handle valid get current user invites params
type GetCurrentUserInvitesHandler interface {
	Handle(GetCurrentUserInvitesParams, interface{}) middleware.Responder
}

// NewGetCurrentUserInvites creates a new http.Handler for the get current user invites operation
func NewGetCurrentUserInvites(ctx *middleware.Context, handler GetCurrentUserInvitesHandler) *GetCurrentUserInvites {
	return &GetCurrentUserInvites{Context: ctx, Handler: handler}
}

/* GetCurrentUserInvites swagger:route GET /users/current/invites current user getCurrentUserInvites

Get group invites of current user

*/
type GetCurrentUserInvites struct {
	Context *middleware.Context
	Handler GetCurrentUserInvitesHandler
}

func (o *GetCurrentUserInvites) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCurrentUserInvitesParams()
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
