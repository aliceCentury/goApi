// Code generated by goa v3.0.4, DO NOT EDIT.
//
// product HTTP server
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	product "calcsvc/gen/product"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the product service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	ProductEndpoint http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the product service endpoints.
func New(
	e *product.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ProductEndpoint", "GET", "/api-product/product/{id}"},
		},
		ProductEndpoint: NewProductEndpointHandler(e.ProductEndpoint, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "product" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ProductEndpoint = m(s.ProductEndpoint)
}

// Mount configures the mux to serve the product endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountProductEndpointHandler(mux, h.ProductEndpoint)
}

// MountProductEndpointHandler configures the mux to serve the "product"
// service "product" endpoint.
func MountProductEndpointHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api-product/product/{id}", f)
}

// NewProductEndpointHandler creates a HTTP handler which loads the HTTP
// request and calls the "product" service "product" endpoint.
func NewProductEndpointHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeProductEndpointRequest(mux, dec)
		encodeResponse = EncodeProductEndpointResponse(enc)
		encodeError    = EncodeProductEndpointError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "product")
		ctx = context.WithValue(ctx, goa.ServiceKey, "product")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
