// Code generated by go-swagger; DO NOT EDIT.

package current_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"mnimidamonbackend/models"
)

// GetCurrentUserGroupsOKCode is the HTTP code returned for type GetCurrentUserGroupsOK
const GetCurrentUserGroupsOKCode int = 200

/*GetCurrentUserGroupsOK Array of groups.

swagger:response getCurrentUserGroupsOK
*/
type GetCurrentUserGroupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Group `json:"body,omitempty"`
}

// NewGetCurrentUserGroupsOK creates GetCurrentUserGroupsOK with default headers values
func NewGetCurrentUserGroupsOK() *GetCurrentUserGroupsOK {

	return &GetCurrentUserGroupsOK{}
}

// WithPayload adds the payload to the get current user groups o k response
func (o *GetCurrentUserGroupsOK) WithPayload(payload []*models.Group) *GetCurrentUserGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user groups o k response
func (o *GetCurrentUserGroupsOK) SetPayload(payload []*models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Group, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCurrentUserGroupsUnauthorizedCode is the HTTP code returned for type GetCurrentUserGroupsUnauthorized
const GetCurrentUserGroupsUnauthorizedCode int = 401

/*GetCurrentUserGroupsUnauthorized Unauthorized.

swagger:response getCurrentUserGroupsUnauthorized
*/
type GetCurrentUserGroupsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCurrentUserGroupsUnauthorized creates GetCurrentUserGroupsUnauthorized with default headers values
func NewGetCurrentUserGroupsUnauthorized() *GetCurrentUserGroupsUnauthorized {

	return &GetCurrentUserGroupsUnauthorized{}
}

// WithPayload adds the payload to the get current user groups unauthorized response
func (o *GetCurrentUserGroupsUnauthorized) WithPayload(payload *models.Error) *GetCurrentUserGroupsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user groups unauthorized response
func (o *GetCurrentUserGroupsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserGroupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
