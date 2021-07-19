// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/adapter/restapi/modelapi"
)

// RegisterComputerOKCode is the HTTP code returned for type RegisterComputerOK
const RegisterComputerOKCode int = 200

/*RegisterComputerOK Computer successfuly created. Returned Computer object and API key that is used in X-COMP-KEY header.

swagger:response registerComputerOK
*/
type RegisterComputerOK struct {

	/*
	  In: Body
	*/
	Payload *modelapi.CreateComputerResponse `json:"body,omitempty"`
}

// NewRegisterComputerOK creates RegisterComputerOK with default headers values
func NewRegisterComputerOK() *RegisterComputerOK {

	return &RegisterComputerOK{}
}

// WithPayload adds the payload to the register computer o k response
func (o *RegisterComputerOK) WithPayload(payload *modelapi.CreateComputerResponse) *RegisterComputerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register computer o k response
func (o *RegisterComputerOK) SetPayload(payload *modelapi.CreateComputerResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterComputerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterComputerBadRequestCode is the HTTP code returned for type RegisterComputerBadRequest
const RegisterComputerBadRequestCode int = 400

/*RegisterComputerBadRequest Supplied parameters were not okay.

swagger:response registerComputerBadRequest
*/
type RegisterComputerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewRegisterComputerBadRequest creates RegisterComputerBadRequest with default headers values
func NewRegisterComputerBadRequest() *RegisterComputerBadRequest {

	return &RegisterComputerBadRequest{}
}

// WithPayload adds the payload to the register computer bad request response
func (o *RegisterComputerBadRequest) WithPayload(payload *modelapi.Error) *RegisterComputerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register computer bad request response
func (o *RegisterComputerBadRequest) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterComputerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterComputerUnauthorizedCode is the HTTP code returned for type RegisterComputerUnauthorized
const RegisterComputerUnauthorizedCode int = 401

/*RegisterComputerUnauthorized Unauthorized.

swagger:response registerComputerUnauthorized
*/
type RegisterComputerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewRegisterComputerUnauthorized creates RegisterComputerUnauthorized with default headers values
func NewRegisterComputerUnauthorized() *RegisterComputerUnauthorized {

	return &RegisterComputerUnauthorized{}
}

// WithPayload adds the payload to the register computer unauthorized response
func (o *RegisterComputerUnauthorized) WithPayload(payload *modelapi.Error) *RegisterComputerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register computer unauthorized response
func (o *RegisterComputerUnauthorized) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterComputerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterComputerInternalServerErrorCode is the HTTP code returned for type RegisterComputerInternalServerError
const RegisterComputerInternalServerErrorCode int = 500

/*RegisterComputerInternalServerError Internal server error.

swagger:response registerComputerInternalServerError
*/
type RegisterComputerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *modelapi.Error `json:"body,omitempty"`
}

// NewRegisterComputerInternalServerError creates RegisterComputerInternalServerError with default headers values
func NewRegisterComputerInternalServerError() *RegisterComputerInternalServerError {

	return &RegisterComputerInternalServerError{}
}

// WithPayload adds the payload to the register computer internal server error response
func (o *RegisterComputerInternalServerError) WithPayload(payload *modelapi.Error) *RegisterComputerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register computer internal server error response
func (o *RegisterComputerInternalServerError) SetPayload(payload *modelapi.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterComputerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
