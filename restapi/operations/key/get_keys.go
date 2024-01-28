// Code generated by go-swagger; DO NOT EDIT.

package key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetKeysHandlerFunc turns a function with the right signature into a get keys handler
type GetKeysHandlerFunc func(GetKeysParams, *AuthPrincipal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetKeysHandlerFunc) Handle(params GetKeysParams, principal *AuthPrincipal) middleware.Responder {
	return fn(params, principal)
}

// GetKeysHandler interface for that can handle valid get keys params
type GetKeysHandler interface {
	Handle(GetKeysParams, *AuthPrincipal) middleware.Responder
}

// NewGetKeys creates a new http.Handler for the get keys operation
func NewGetKeys(ctx *middleware.Context, handler GetKeysHandler) *GetKeys {
	return &GetKeys{Context: ctx, Handler: handler}
}

/*
	GetKeys swagger:route GET /key key getKeys

Get filtered key list
*/
type GetKeys struct {
	Context *middleware.Context
	Handler GetKeysHandler
}

func (o *GetKeys) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetKeysParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *AuthPrincipal
	if uprinc != nil {
		principal = uprinc.(*AuthPrincipal) // this is really a AuthPrincipal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}