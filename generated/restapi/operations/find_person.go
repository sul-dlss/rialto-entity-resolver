// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindPersonHandlerFunc turns a function with the right signature into a find person handler
type FindPersonHandlerFunc func(FindPersonParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn FindPersonHandlerFunc) Handle(params FindPersonParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// FindPersonHandler interface for that can handle valid find person params
type FindPersonHandler interface {
	Handle(FindPersonParams, interface{}) middleware.Responder
}

// NewFindPerson creates a new http.Handler for the find person operation
func NewFindPerson(ctx *middleware.Context, handler FindPersonHandler) *FindPerson {
	return &FindPerson{Context: ctx, Handler: handler}
}

/*FindPerson swagger:route GET /person findPerson

FindPerson find person API

*/
type FindPerson struct {
	Context *middleware.Context
	Handler FindPersonHandler
}

func (o *FindPerson) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindPersonParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
