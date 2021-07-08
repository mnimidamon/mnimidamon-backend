// Code generated by go-swagger; DO NOT EDIT.

package invite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AcceptCurrentUserInviteHandlerFunc turns a function with the right signature into a accept current user invite handler
type AcceptCurrentUserInviteHandlerFunc func(AcceptCurrentUserInviteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AcceptCurrentUserInviteHandlerFunc) Handle(params AcceptCurrentUserInviteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AcceptCurrentUserInviteHandler interface for that can handle valid accept current user invite params
type AcceptCurrentUserInviteHandler interface {
	Handle(AcceptCurrentUserInviteParams, interface{}) middleware.Responder
}

// NewAcceptCurrentUserInvite creates a new http.Handler for the accept current user invite operation
func NewAcceptCurrentUserInvite(ctx *middleware.Context, handler AcceptCurrentUserInviteHandler) *AcceptCurrentUserInvite {
	return &AcceptCurrentUserInvite{Context: ctx, Handler: handler}
}

/* AcceptCurrentUserInvite swagger:route POST /users/current/invites/{group_id}/accept invite acceptCurrentUserInvite

Accept a group invite

*/
type AcceptCurrentUserInvite struct {
	Context *middleware.Context
	Handler AcceptCurrentUserInviteHandler
}

func (o *AcceptCurrentUserInvite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAcceptCurrentUserInviteParams()
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
