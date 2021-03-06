// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/adapter/restapi/modelapi"
)

// GetCurrentUserGroupComputersOKCode is the HTTP code returned for type GetCurrentUserGroupComputersOK
const GetCurrentUserGroupComputersOKCode int = 200

/*GetCurrentUserGroupComputersOK Array of the denoted group computers.

swagger:response getCurrentUserGroupComputersOK
*/
type GetCurrentUserGroupComputersOK struct {

	/*
	  In: Body
	*/
	Payload []*modelapi.GroupComputer `json:"body,omitempty"`
}

// NewGetCurrentUserGroupComputersOK creates GetCurrentUserGroupComputersOK with default headers values
func NewGetCurrentUserGroupComputersOK() *GetCurrentUserGroupComputersOK {

	return &GetCurrentUserGroupComputersOK{}
}

// WithPayload adds the payload to the get current user group computers o k response
func (o *GetCurrentUserGroupComputersOK) WithPayload(payload []*modelapi.GroupComputer) *GetCurrentUserGroupComputersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user group computers o k response
func (o *GetCurrentUserGroupComputersOK) SetPayload(payload []*modelapi.GroupComputer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserGroupComputersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*modelapi.GroupComputer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCurrentUserGroupComputersUnauthorizedCode is the HTTP code returned for type GetCurrentUserGroupComputersUnauthorized
const GetCurrentUserGroupComputersUnauthorizedCode int = 401

/*GetCurrentUserGroupComputersUnauthorized Unauthorized.

swagger:response getCurrentUserGroupComputersUnauthorized
*/
type GetCurrentUserGroupComputersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewGetCurrentUserGroupComputersUnauthorized creates GetCurrentUserGroupComputersUnauthorized with default headers values
func NewGetCurrentUserGroupComputersUnauthorized() *GetCurrentUserGroupComputersUnauthorized {

	return &GetCurrentUserGroupComputersUnauthorized{}
}

// WithPayload adds the payload to the get current user group computers unauthorized response
func (o *GetCurrentUserGroupComputersUnauthorized) WithPayload(payload *modelapi.Error) *GetCurrentUserGroupComputersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user group computers unauthorized response
func (o *GetCurrentUserGroupComputersUnauthorized) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserGroupComputersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentUserGroupComputersInternalServerErrorCode is the HTTP code returned for type GetCurrentUserGroupComputersInternalServerError
const GetCurrentUserGroupComputersInternalServerErrorCode int = 500

/*GetCurrentUserGroupComputersInternalServerError Internal server error.

swagger:response getCurrentUserGroupComputersInternalServerError
*/
type GetCurrentUserGroupComputersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewGetCurrentUserGroupComputersInternalServerError creates GetCurrentUserGroupComputersInternalServerError with default headers values
func NewGetCurrentUserGroupComputersInternalServerError() *GetCurrentUserGroupComputersInternalServerError {

	return &GetCurrentUserGroupComputersInternalServerError{}
}

// WithPayload adds the payload to the get current user group computers internal server error response
func (o *GetCurrentUserGroupComputersInternalServerError) WithPayload(payload *modelapi.Error) *GetCurrentUserGroupComputersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user group computers internal server error response
func (o *GetCurrentUserGroupComputersInternalServerError) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserGroupComputersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
