// Code generated by go-swagger; DO NOT EDIT.

package invite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/adapter/restapi/modelapi"
)

// AcceptCurrentUserInviteOKCode is the HTTP code returned for type AcceptCurrentUserInviteOK
const AcceptCurrentUserInviteOKCode int = 200

/*AcceptCurrentUserInviteOK The accepted group object.

swagger:response acceptCurrentUserInviteOK
*/
type AcceptCurrentUserInviteOK struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Group `json:"body,omitempty"`
}

// NewAcceptCurrentUserInviteOK creates AcceptCurrentUserInviteOK with default headers values
func NewAcceptCurrentUserInviteOK() *AcceptCurrentUserInviteOK {

	return &AcceptCurrentUserInviteOK{}
}

// WithPayload adds the payload to the accept current user invite o k response
func (o *AcceptCurrentUserInviteOK) WithPayload(payload *modelapi.Group) *AcceptCurrentUserInviteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the accept current user invite o k response
func (o *AcceptCurrentUserInviteOK) SetPayload(payload *modelapi.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcceptCurrentUserInviteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcceptCurrentUserInviteUnauthorizedCode is the HTTP code returned for type AcceptCurrentUserInviteUnauthorized
const AcceptCurrentUserInviteUnauthorizedCode int = 401

/*AcceptCurrentUserInviteUnauthorized Unauthorized.

swagger:response acceptCurrentUserInviteUnauthorized
*/
type AcceptCurrentUserInviteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewAcceptCurrentUserInviteUnauthorized creates AcceptCurrentUserInviteUnauthorized with default headers values
func NewAcceptCurrentUserInviteUnauthorized() *AcceptCurrentUserInviteUnauthorized {

	return &AcceptCurrentUserInviteUnauthorized{}
}

// WithPayload adds the payload to the accept current user invite unauthorized response
func (o *AcceptCurrentUserInviteUnauthorized) WithPayload(payload *modelapi.Error) *AcceptCurrentUserInviteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the accept current user invite unauthorized response
func (o *AcceptCurrentUserInviteUnauthorized) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcceptCurrentUserInviteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcceptCurrentUserInviteInternalServerErrorCode is the HTTP code returned for type AcceptCurrentUserInviteInternalServerError
const AcceptCurrentUserInviteInternalServerErrorCode int = 500

/*AcceptCurrentUserInviteInternalServerError Internal server error.

swagger:response acceptCurrentUserInviteInternalServerError
*/
type AcceptCurrentUserInviteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewAcceptCurrentUserInviteInternalServerError creates AcceptCurrentUserInviteInternalServerError with default headers values
func NewAcceptCurrentUserInviteInternalServerError() *AcceptCurrentUserInviteInternalServerError {

	return &AcceptCurrentUserInviteInternalServerError{}
}

// WithPayload adds the payload to the accept current user invite internal server error response
func (o *AcceptCurrentUserInviteInternalServerError) WithPayload(payload *modelapi.Error) *AcceptCurrentUserInviteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the accept current user invite internal server error response
func (o *AcceptCurrentUserInviteInternalServerError) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcceptCurrentUserInviteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
