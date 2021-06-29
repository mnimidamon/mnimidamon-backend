// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetGroupParams creates a new GetGroupParams object
//
// There are no default values defined in the spec.
func NewGetGroupParams() GetGroupParams {

	return GetGroupParams{}
}

// GetGroupParams contains all the bound params for the get group operation
// typically these are obtained from a http.Request
//
// swagger:parameters getGroup
type GetGroupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Numeric ID of the Group.
	  Required: true
	  In: path
	*/
	GroupID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetGroupParams() beforehand.
func (o *GetGroupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGroupID, rhkGroupID, _ := route.Params.GetOK("group_id")
	if err := o.bindGroupID(rGroupID, rhkGroupID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGroupID binds and validates parameter GroupID from path.
func (o *GetGroupParams) bindGroupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("group_id", "path", "int64", raw)
	}
	o.GroupID = value

	return nil
}
