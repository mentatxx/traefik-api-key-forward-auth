// Code generated by go-swagger; DO NOT EDIT.

package key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteKeyOKCode is the HTTP code returned for type DeleteKeyOK
const DeleteKeyOKCode int = 200

/*
DeleteKeyOK successful operation

swagger:response deleteKeyOK
*/
type DeleteKeyOK struct {
}

// NewDeleteKeyOK creates DeleteKeyOK with default headers values
func NewDeleteKeyOK() *DeleteKeyOK {

	return &DeleteKeyOK{}
}

// WriteResponse to the client
func (o *DeleteKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteKeyBadRequestCode is the HTTP code returned for type DeleteKeyBadRequest
const DeleteKeyBadRequestCode int = 400

/*
DeleteKeyBadRequest Invalid request supplied

swagger:response deleteKeyBadRequest
*/
type DeleteKeyBadRequest struct {
}

// NewDeleteKeyBadRequest creates DeleteKeyBadRequest with default headers values
func NewDeleteKeyBadRequest() *DeleteKeyBadRequest {

	return &DeleteKeyBadRequest{}
}

// WriteResponse to the client
func (o *DeleteKeyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}