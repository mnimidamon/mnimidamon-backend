// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/adapter/restapi/modelapi"
)

// GetGroupBackupsOKCode is the HTTP code returned for type GetGroupBackupsOK
const GetGroupBackupsOKCode int = 200

/*GetGroupBackupsOK Array of the group backups.

swagger:response getGroupBackupsOK
*/
type GetGroupBackupsOK struct {

	/*
	  In: Body
	*/
	Payload []*modelapi.Backup `json:"body,omitempty"`
}

// NewGetGroupBackupsOK creates GetGroupBackupsOK with default headers values
func NewGetGroupBackupsOK() *GetGroupBackupsOK {

	return &GetGroupBackupsOK{}
}

// WithPayload adds the payload to the get group backups o k response
func (o *GetGroupBackupsOK) WithPayload(payload []*modelapi.Backup) *GetGroupBackupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group backups o k response
func (o *GetGroupBackupsOK) SetPayload(payload []*modelapi.Backup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupBackupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*modelapi.Backup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetGroupBackupsUnauthorizedCode is the HTTP code returned for type GetGroupBackupsUnauthorized
const GetGroupBackupsUnauthorizedCode int = 401

/*GetGroupBackupsUnauthorized Unauthorized.

swagger:response getGroupBackupsUnauthorized
*/
type GetGroupBackupsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewGetGroupBackupsUnauthorized creates GetGroupBackupsUnauthorized with default headers values
func NewGetGroupBackupsUnauthorized() *GetGroupBackupsUnauthorized {

	return &GetGroupBackupsUnauthorized{}
}

// WithPayload adds the payload to the get group backups unauthorized response
func (o *GetGroupBackupsUnauthorized) WithPayload(payload *modelapi.Error) *GetGroupBackupsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group backups unauthorized response
func (o *GetGroupBackupsUnauthorized) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupBackupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupBackupsInternalServerErrorCode is the HTTP code returned for type GetGroupBackupsInternalServerError
const GetGroupBackupsInternalServerErrorCode int = 500

/*GetGroupBackupsInternalServerError Internal server error.

swagger:response getGroupBackupsInternalServerError
*/
type GetGroupBackupsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewGetGroupBackupsInternalServerError creates GetGroupBackupsInternalServerError with default headers values
func NewGetGroupBackupsInternalServerError() *GetGroupBackupsInternalServerError {

	return &GetGroupBackupsInternalServerError{}
}

// WithPayload adds the payload to the get group backups internal server error response
func (o *GetGroupBackupsInternalServerError) WithPayload(payload *modelapi.Error) *GetGroupBackupsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group backups internal server error response
func (o *GetGroupBackupsInternalServerError) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupBackupsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
