// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ridhodinar/dans-tes/models"
)

// GetAllJobsOKCode is the HTTP code returned for type GetAllJobsOK
const GetAllJobsOKCode int = 200

/*
GetAllJobsOK All Jobs

swagger:response getAllJobsOK
*/
type GetAllJobsOK struct {

	/*
	  In: Body
	*/
	Payload models.JobList `json:"body,omitempty"`
}

// NewGetAllJobsOK creates GetAllJobsOK with default headers values
func NewGetAllJobsOK() *GetAllJobsOK {

	return &GetAllJobsOK{}
}

// WithPayload adds the payload to the get all jobs o k response
func (o *GetAllJobsOK) WithPayload(payload models.JobList) *GetAllJobsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all jobs o k response
func (o *GetAllJobsOK) SetPayload(payload models.JobList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllJobsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.JobList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllJobsBadRequestCode is the HTTP code returned for type GetAllJobsBadRequest
const GetAllJobsBadRequestCode int = 400

/*
GetAllJobsBadRequest Bad Request

swagger:response getAllJobsBadRequest
*/
type GetAllJobsBadRequest struct {
}

// NewGetAllJobsBadRequest creates GetAllJobsBadRequest with default headers values
func NewGetAllJobsBadRequest() *GetAllJobsBadRequest {

	return &GetAllJobsBadRequest{}
}

// WriteResponse to the client
func (o *GetAllJobsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAllJobsNotFoundCode is the HTTP code returned for type GetAllJobsNotFound
const GetAllJobsNotFoundCode int = 404

/*
GetAllJobsNotFound Jobs Not Found

swagger:response getAllJobsNotFound
*/
type GetAllJobsNotFound struct {
}

// NewGetAllJobsNotFound creates GetAllJobsNotFound with default headers values
func NewGetAllJobsNotFound() *GetAllJobsNotFound {

	return &GetAllJobsNotFound{}
}

// WriteResponse to the client
func (o *GetAllJobsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetAllJobsInternalServerErrorCode is the HTTP code returned for type GetAllJobsInternalServerError
const GetAllJobsInternalServerErrorCode int = 500

/*
GetAllJobsInternalServerError Server error

swagger:response getAllJobsInternalServerError
*/
type GetAllJobsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAllJobsInternalServerError creates GetAllJobsInternalServerError with default headers values
func NewGetAllJobsInternalServerError() *GetAllJobsInternalServerError {

	return &GetAllJobsInternalServerError{}
}

// WithPayload adds the payload to the get all jobs internal server error response
func (o *GetAllJobsInternalServerError) WithPayload(payload string) *GetAllJobsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all jobs internal server error response
func (o *GetAllJobsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllJobsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
