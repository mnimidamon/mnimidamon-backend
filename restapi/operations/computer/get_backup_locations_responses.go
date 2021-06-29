// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/models"
)

// GetBackupLocationsOKCode is the HTTP code returned for type GetBackupLocationsOK
const GetBackupLocationsOKCode int = 200

/*GetBackupLocationsOK Users and their computers that have the backup

swagger:response getBackupLocationsOK
*/
type GetBackupLocationsOK struct {

	/*
	  In: Body
	*/
	Payload models.StoredBackupsComputerResponse `json:"body,omitempty"`
}

// NewGetBackupLocationsOK creates GetBackupLocationsOK with default headers values
func NewGetBackupLocationsOK() *GetBackupLocationsOK {

	return &GetBackupLocationsOK{}
}

// WithPayload adds the payload to the get backup locations o k response
func (o *GetBackupLocationsOK) WithPayload(payload models.StoredBackupsComputerResponse) *GetBackupLocationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backup locations o k response
func (o *GetBackupLocationsOK) SetPayload(payload models.StoredBackupsComputerResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackupLocationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.StoredBackupsComputerResponse{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBackupLocationsUnauthorizedCode is the HTTP code returned for type GetBackupLocationsUnauthorized
const GetBackupLocationsUnauthorizedCode int = 401

/*GetBackupLocationsUnauthorized Unauthorized.

swagger:response getBackupLocationsUnauthorized
*/
type GetBackupLocationsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackupLocationsUnauthorized creates GetBackupLocationsUnauthorized with default headers values
func NewGetBackupLocationsUnauthorized() *GetBackupLocationsUnauthorized {

	return &GetBackupLocationsUnauthorized{}
}

// WithPayload adds the payload to the get backup locations unauthorized response
func (o *GetBackupLocationsUnauthorized) WithPayload(payload *models.Error) *GetBackupLocationsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backup locations unauthorized response
func (o *GetBackupLocationsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackupLocationsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackupLocationsNotFoundCode is the HTTP code returned for type GetBackupLocationsNotFound
const GetBackupLocationsNotFoundCode int = 404

/*GetBackupLocationsNotFound The specified resource was not found.

swagger:response getBackupLocationsNotFound
*/
type GetBackupLocationsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackupLocationsNotFound creates GetBackupLocationsNotFound with default headers values
func NewGetBackupLocationsNotFound() *GetBackupLocationsNotFound {

	return &GetBackupLocationsNotFound{}
}

// WithPayload adds the payload to the get backup locations not found response
func (o *GetBackupLocationsNotFound) WithPayload(payload *models.Error) *GetBackupLocationsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backup locations not found response
func (o *GetBackupLocationsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackupLocationsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
