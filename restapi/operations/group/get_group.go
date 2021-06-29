// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGroupHandlerFunc turns a function with the right signature into a get group handler
type GetGroupHandlerFunc func(GetGroupParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGroupHandlerFunc) Handle(params GetGroupParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetGroupHandler interface for that can handle valid get group params
type GetGroupHandler interface {
	Handle(GetGroupParams, interface{}) middleware.Responder
}

// NewGetGroup creates a new http.Handler for the get group operation
func NewGetGroup(ctx *middleware.Context, handler GetGroupHandler) *GetGroup {
	return &GetGroup{Context: ctx, Handler: handler}
}

/* GetGroup swagger:route GET /users/current/groups/{group_id} group getGroup

Get a group

*/
type GetGroup struct {
	Context *middleware.Context
	Handler GetGroupHandler
}

func (o *GetGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGroupParams()
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
