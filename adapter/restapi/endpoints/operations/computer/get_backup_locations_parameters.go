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

// NewGetBackupLocationsParams creates a new GetBackupLocationsParams object
//
// There are no default values defined in the spec.
func NewGetBackupLocationsParams() GetBackupLocationsParams {

	return GetBackupLocationsParams{}
}

// GetBackupLocationsParams contains all the bound params for the get backup locations operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBackupLocations
type GetBackupLocationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Numeric ID of the Backup.
	  Required: true
	  In: path
	*/
	BackupID int64
	/*Numeric ID of the Group.
	  Required: true
	  In: path
	*/
	GroupID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBackupLocationsParams() beforehand.
func (o *GetBackupLocationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBackupID, rhkBackupID, _ := route.Params.GetOK("backup_id")
	if err := o.bindBackupID(rBackupID, rhkBackupID, route.Formats); err != nil {
		res = append(res, err)
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

// bindBackupID binds and validates parameter BackupID from path.
func (o *GetBackupLocationsParams) bindBackupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("backup_id", "path", "int64", raw)
	}
	o.BackupID = value

	return nil
}

// bindGroupID binds and validates parameter GroupID from path.
func (o *GetBackupLocationsParams) bindGroupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
