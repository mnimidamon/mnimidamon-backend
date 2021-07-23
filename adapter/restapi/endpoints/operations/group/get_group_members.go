// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGroupMembersHandlerFunc turns a function with the right signature into a get group members handler
type GetGroupMembersHandlerFunc func(GetGroupMembersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGroupMembersHandlerFunc) Handle(params GetGroupMembersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetGroupMembersHandler interface for that can handle valid get group members params
type GetGroupMembersHandler interface {
	Handle(GetGroupMembersParams, interface{}) middleware.Responder
}

// NewGetGroupMembers creates a new http.Handler for the get group members operation
func NewGetGroupMembers(ctx *middleware.Context, handler GetGroupMembersHandler) *GetGroupMembers {
	return &GetGroupMembers{Context: ctx, Handler: handler}
}

/* GetGroupMembers swagger:route GET /users/current/groups/{group_id}/members group getGroupMembers

Get group members

*/
type GetGroupMembers struct {
	Context *middleware.Context
	Handler GetGroupMembersHandler
}

func (o *GetGroupMembers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGroupMembersParams()
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