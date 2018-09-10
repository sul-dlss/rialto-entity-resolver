// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindOrCreatePersonHandlerFunc turns a function with the right signature into a find or create person handler
type FindOrCreatePersonHandlerFunc func(FindOrCreatePersonParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindOrCreatePersonHandlerFunc) Handle(params FindOrCreatePersonParams) middleware.Responder {
	return fn(params)
}

// FindOrCreatePersonHandler interface for that can handle valid find or create person params
type FindOrCreatePersonHandler interface {
	Handle(FindOrCreatePersonParams) middleware.Responder
}

// NewFindOrCreatePerson creates a new http.Handler for the find or create person operation
func NewFindOrCreatePerson(ctx *middleware.Context, handler FindOrCreatePersonHandler) *FindOrCreatePerson {
	return &FindOrCreatePerson{Context: ctx, Handler: handler}
}

/*FindOrCreatePerson swagger:route GET /person findOrCreatePerson

FindOrCreatePerson find or create person API

*/
type FindOrCreatePerson struct {
	Context *middleware.Context
	Handler FindOrCreatePersonHandler
}

func (o *FindOrCreatePerson) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindOrCreatePersonParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
