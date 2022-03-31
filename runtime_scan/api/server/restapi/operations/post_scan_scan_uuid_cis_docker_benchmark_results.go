// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostScanScanUUIDCisDockerBenchmarkResultsHandlerFunc turns a function with the right signature into a post scan scan UUID cis docker benchmark results handler
type PostScanScanUUIDCisDockerBenchmarkResultsHandlerFunc func(PostScanScanUUIDCisDockerBenchmarkResultsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostScanScanUUIDCisDockerBenchmarkResultsHandlerFunc) Handle(params PostScanScanUUIDCisDockerBenchmarkResultsParams) middleware.Responder {
	return fn(params)
}

// PostScanScanUUIDCisDockerBenchmarkResultsHandler interface for that can handle valid post scan scan UUID cis docker benchmark results params
type PostScanScanUUIDCisDockerBenchmarkResultsHandler interface {
	Handle(PostScanScanUUIDCisDockerBenchmarkResultsParams) middleware.Responder
}

// NewPostScanScanUUIDCisDockerBenchmarkResults creates a new http.Handler for the post scan scan UUID cis docker benchmark results operation
func NewPostScanScanUUIDCisDockerBenchmarkResults(ctx *middleware.Context, handler PostScanScanUUIDCisDockerBenchmarkResultsHandler) *PostScanScanUUIDCisDockerBenchmarkResults {
	return &PostScanScanUUIDCisDockerBenchmarkResults{Context: ctx, Handler: handler}
}

/* PostScanScanUUIDCisDockerBenchmarkResults swagger:route POST /scan/{scan-uuid}/cisDockerBenchmark/results postScanScanUuidCisDockerBenchmarkResults

Report an image CIS docker benchmark scan for a specific scan UUID.

*/
type PostScanScanUUIDCisDockerBenchmarkResults struct {
	Context *middleware.Context
	Handler PostScanScanUUIDCisDockerBenchmarkResultsHandler
}

func (o *PostScanScanUUIDCisDockerBenchmarkResults) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostScanScanUUIDCisDockerBenchmarkResultsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}