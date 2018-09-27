// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindOneHandlerFunc turns a function with the right signature into a find one handler
type FindOneHandlerFunc func(FindOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindOneHandlerFunc) Handle(params FindOneParams) middleware.Responder {
	return fn(params)
}

// FindOneHandler interface for that can handle valid find one params
type FindOneHandler interface {
	Handle(FindOneParams) middleware.Responder
}

// NewFindOne creates a new http.Handler for the find one operation
func NewFindOne(ctx *middleware.Context, handler FindOneHandler) *FindOne {
	return &FindOne{Context: ctx, Handler: handler}
}

/*FindOne swagger:route GET /{id} todos findOne

FindOne find one API

*/
type FindOne struct {
	Context *middleware.Context
	Handler FindOneHandler
}

func (o *FindOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}