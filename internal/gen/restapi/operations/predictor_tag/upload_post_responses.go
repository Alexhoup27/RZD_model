// Code generated by go-swagger; DO NOT EDIT.

package predictor_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"service/internal/gen/models"
)

// UploadPostOKCode is the HTTP code returned for type UploadPostOK
const UploadPostOKCode int = 200

/*
UploadPostOK Возвращаемые ответы

swagger:response uploadPostOK
*/
type UploadPostOK struct {

	/*
	  In: Body
	*/
	Payload *UploadPostOKBody `json:"body,omitempty"`
}

// NewUploadPostOK creates UploadPostOK with default headers values
func NewUploadPostOK() *UploadPostOK {

	return &UploadPostOK{}
}

// WithPayload adds the payload to the upload post o k response
func (o *UploadPostOK) WithPayload(payload *UploadPostOKBody) *UploadPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload post o k response
func (o *UploadPostOK) SetPayload(payload *UploadPostOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadPostInternalServerErrorCode is the HTTP code returned for type UploadPostInternalServerError
const UploadPostInternalServerErrorCode int = 500

/*
UploadPostInternalServerError Ошибка сервера

swagger:response uploadPostInternalServerError
*/
type UploadPostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewUploadPostInternalServerError creates UploadPostInternalServerError with default headers values
func NewUploadPostInternalServerError() *UploadPostInternalServerError {

	return &UploadPostInternalServerError{}
}

// WithPayload adds the payload to the upload post internal server error response
func (o *UploadPostInternalServerError) WithPayload(payload *models.Error500) *UploadPostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload post internal server error response
func (o *UploadPostInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadPostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}