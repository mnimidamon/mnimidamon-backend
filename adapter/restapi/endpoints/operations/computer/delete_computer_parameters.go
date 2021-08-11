// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteComputerParams creates a new DeleteComputerParams object
//
// There are no default values defined in the spec.
func NewDeleteComputerParams() DeleteComputerParams {

	return DeleteComputerParams{}
}

// DeleteComputerParams contains all the bound params for the delete computer operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteComputer
type DeleteComputerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Numeric ID of the Computer.
	  Required: true
	  In: path
	*/
	ComputerID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteComputerParams() beforehand.
func (o *DeleteComputerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rComputerID, rhkComputerID, _ := route.Params.GetOK("computer_id")
	if err := o.bindComputerID(rComputerID, rhkComputerID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindComputerID binds and validates parameter ComputerID from path.
func (o *DeleteComputerParams) bindComputerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("computer_id", "path", "int64", raw)
	}
	o.ComputerID = value

	return nil
}
