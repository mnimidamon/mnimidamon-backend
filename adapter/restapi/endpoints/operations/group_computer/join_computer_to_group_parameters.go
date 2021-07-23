// Code generated by go-swagger; DO NOT EDIT.

package group_computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"mnimidamonbackend/adapter/restapi/modelapi"
)

// NewJoinComputerToGroupParams creates a new JoinComputerToGroupParams object
//
// There are no default values defined in the spec.
func NewJoinComputerToGroupParams() JoinComputerToGroupParams {

	return JoinComputerToGroupParams{}
}

// JoinComputerToGroupParams contains all the bound params for the join computer to group operation
// typically these are obtained from a http.Request
//
// swagger:parameters joinComputerToGroup
type JoinComputerToGroupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Group creation payload.
	  Required: true
	  In: body
	*/
	Body *modelapi.CreateGroupComputerPayload
	/*Numeric ID of the Group.
	  Required: true
	  In: path
	*/
	GroupID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewJoinComputerToGroupParams() beforehand.
func (o *JoinComputerToGroupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body modelapi.CreateGroupComputerPayload
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

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
func (o *JoinComputerToGroupParams) bindGroupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
