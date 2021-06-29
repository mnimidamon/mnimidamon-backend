// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/models"
)

// RegisterUserOKCode is the HTTP code returned for type RegisterUserOK
const RegisterUserOKCode int = 200

/*RegisterUserOK Authorization successful. Access token and created user response.

swagger:response registerUserOK
*/
type RegisterUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.RegisterResponse `json:"body,omitempty"`
}

// NewRegisterUserOK creates RegisterUserOK with default headers values
func NewRegisterUserOK() *RegisterUserOK {

	return &RegisterUserOK{}
}

// WithPayload adds the payload to the register user o k response
func (o *RegisterUserOK) WithPayload(payload *models.RegisterResponse) *RegisterUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user o k response
func (o *RegisterUserOK) SetPayload(payload *models.RegisterResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserNotFoundCode is the HTTP code returned for type RegisterUserNotFound
const RegisterUserNotFoundCode int = 404

/*RegisterUserNotFound Supplied parameters were not okay.

swagger:response registerUserNotFound
*/
type RegisterUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterUserNotFound creates RegisterUserNotFound with default headers values
func NewRegisterUserNotFound() *RegisterUserNotFound {

	return &RegisterUserNotFound{}
}

// WithPayload adds the payload to the register user not found response
func (o *RegisterUserNotFound) WithPayload(payload *models.Error) *RegisterUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user not found response
func (o *RegisterUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
