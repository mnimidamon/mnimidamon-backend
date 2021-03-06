// Code generated by go-swagger; DO NOT EDIT.

package group_computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGroupComputersOfComputerHandlerFunc turns a function with the right signature into a get group computers of computer handler
type GetGroupComputersOfComputerHandlerFunc func(GetGroupComputersOfComputerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGroupComputersOfComputerHandlerFunc) Handle(params GetGroupComputersOfComputerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetGroupComputersOfComputerHandler interface for that can handle valid get group computers of computer params
type GetGroupComputersOfComputerHandler interface {
	Handle(GetGroupComputersOfComputerParams, interface{}) middleware.Responder
}

// NewGetGroupComputersOfComputer creates a new http.Handler for the get group computers of computer operation
func NewGetGroupComputersOfComputer(ctx *middleware.Context, handler GetGroupComputersOfComputerHandler) *GetGroupComputersOfComputer {
	return &GetGroupComputersOfComputer{Context: ctx, Handler: handler}
}

/* GetGroupComputersOfComputer swagger:route GET /users/current/computers/{computer_id}/groups group computer getGroupComputersOfComputer

Get group computers of computer

*/
type GetGroupComputersOfComputer struct {
	Context *middleware.Context
	Handler GetGroupComputersOfComputerHandler
}

func (o *GetGroupComputersOfComputer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGroupComputersOfComputerParams()
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
