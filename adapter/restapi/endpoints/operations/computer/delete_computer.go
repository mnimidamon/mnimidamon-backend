// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteComputerHandlerFunc turns a function with the right signature into a delete computer handler
type DeleteComputerHandlerFunc func(DeleteComputerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteComputerHandlerFunc) Handle(params DeleteComputerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteComputerHandler interface for that can handle valid delete computer params
type DeleteComputerHandler interface {
	Handle(DeleteComputerParams, interface{}) middleware.Responder
}

// NewDeleteComputer creates a new http.Handler for the delete computer operation
func NewDeleteComputer(ctx *middleware.Context, handler DeleteComputerHandler) *DeleteComputer {
	return &DeleteComputer{Context: ctx, Handler: handler}
}

/* DeleteComputer swagger:route DELETE /users/current/computers/{computer_id} computer deleteComputer

Delete a computer

Deletes the computer with its group memberships

*/
type DeleteComputer struct {
	Context *middleware.Context
	Handler DeleteComputerHandler
}

func (o *DeleteComputer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteComputerParams()
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
