package authentication

import (
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"net/http"
)

var (
	ErrInvalidComputerAuthToken    = errors.New("ErrInvalidComputerAuthToken")
	ErrInvalidUserAuthToken        = errors.New("ErrInvalidUserAuthToken")
	ErrExtractingUserAuthToken     = errors.New("ErrExtractingUserAuthToken")
	ErrExtractingComputerAuthToken = errors.New("ErrExtractingComputerAuthToken")
)

type unauthorizedErrorResponder struct {
	Payload *modelapi.Error `json:"body,omitempty"`
}

func newUnauthorizedErrorResponder(err error) middleware.Responder {
	if err == nil {
		return &unauthorizedErrorResponder{}
	}
	return &unauthorizedErrorResponder{
		Payload: &modelapi.Error{
			Message: err.Error(),
		},
	}
}

// This is to suffice the middleware.Responder
func (o *unauthorizedErrorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

type internalServerErrorResponder struct {
	Payload *modelapi.Error `json:"body,omitempty"`
}

func newInternalServerErrorResponder(err error) middleware.Responder {
	if err == nil {
		return &internalServerErrorResponder{}
	}
	return &internalServerErrorResponder{
		Payload: &modelapi.Error{
			Message: err.Error(),
		},
	}
}

// This is to suffice the middleware.Responder
func (o *internalServerErrorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

type badRequestErrorResponder struct {
	Payload *modelapi.Error `json:"body,omitempty"`
}

func newBadRequestErrorResponder(err error) middleware.Responder {
	if err == nil {
		return &badRequestErrorResponder{}
	}
	return &internalServerErrorResponder{
		Payload: &modelapi.Error{
			Message: err.Error(),
		},
	}
}

// This is to suffice the middleware.Responder
func (o *badRequestErrorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
