// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// InitializeGroupBackupDeletionURL generates an URL for the initialize group backup deletion operation
type InitializeGroupBackupDeletionURL struct {
	BackupID int64
	GroupID  int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *InitializeGroupBackupDeletionURL) WithBasePath(bp string) *InitializeGroupBackupDeletionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *InitializeGroupBackupDeletionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *InitializeGroupBackupDeletionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/users/current/computers/current/groups/{group_id}/backups/{backup_id}"

	backupID := swag.FormatInt64(o.BackupID)
	if backupID != "" {
		_path = strings.Replace(_path, "{backup_id}", backupID, -1)
	} else {
		return nil, errors.New("backupId is required on InitializeGroupBackupDeletionURL")
	}

	groupID := swag.FormatInt64(o.GroupID)
	if groupID != "" {
		_path = strings.Replace(_path, "{group_id}", groupID, -1)
	} else {
		return nil, errors.New("groupId is required on InitializeGroupBackupDeletionURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *InitializeGroupBackupDeletionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *InitializeGroupBackupDeletionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *InitializeGroupBackupDeletionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on InitializeGroupBackupDeletionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on InitializeGroupBackupDeletionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *InitializeGroupBackupDeletionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
