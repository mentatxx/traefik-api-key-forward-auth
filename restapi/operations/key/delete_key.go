// Code generated by go-swagger; DO NOT EDIT.

package key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/mentatxx/traefik-api-key-forward-auth/models"
)

// DeleteKeyHandlerFunc turns a function with the right signature into a delete key handler
type DeleteKeyHandlerFunc func(DeleteKeyParams, *models.AuthPrincipal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteKeyHandlerFunc) Handle(params DeleteKeyParams, principal *models.AuthPrincipal) middleware.Responder {
	return fn(params, principal)
}

// DeleteKeyHandler interface for that can handle valid delete key params
type DeleteKeyHandler interface {
	Handle(DeleteKeyParams, *models.AuthPrincipal) middleware.Responder
}

// NewDeleteKey creates a new http.Handler for the delete key operation
func NewDeleteKey(ctx *middleware.Context, handler DeleteKeyHandler) *DeleteKey {
	return &DeleteKey{Context: ctx, Handler: handler}
}

/*
	DeleteKey swagger:route DELETE /key/{id} key deleteKey

Delete key by ID
*/
type DeleteKey struct {
	Context *middleware.Context
	Handler DeleteKeyHandler
}

func (o *DeleteKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteKeyParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.AuthPrincipal
	if uprinc != nil {
		principal = uprinc.(*models.AuthPrincipal) // this is really a models.AuthPrincipal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
