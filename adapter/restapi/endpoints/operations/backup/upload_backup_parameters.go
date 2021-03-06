// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UploadBackupMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var UploadBackupMaxParseMemory int64 = 32 << 20

// NewUploadBackupParams creates a new UploadBackupParams object
//
// There are no default values defined in the spec.
func NewUploadBackupParams() UploadBackupParams {

	return UploadBackupParams{}
}

// UploadBackupParams contains all the bound params for the upload backup operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadBackup
type UploadBackupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The encodec backup file.
	  In: formData
	*/
	BackupData io.ReadCloser
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
// To ensure default values, the struct must have been initialized with NewUploadBackupParams() beforehand.
func (o *UploadBackupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(UploadBackupMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	backupData, backupDataHeader, err := r.FormFile("backup_data")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "backupData", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindBackupData(backupData, backupDataHeader); err != nil {
		res = append(res, err)
	} else {
		o.BackupData = &runtime.File{Data: backupData, Header: backupDataHeader}
	}

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

// bindBackupData binds file parameter BackupData.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadBackupParams) bindBackupData(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindBackupID binds and validates parameter BackupID from path.
func (o *UploadBackupParams) bindBackupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UploadBackupParams) bindGroupID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
