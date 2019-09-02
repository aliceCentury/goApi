// Code generated by goa v3.0.4, DO NOT EDIT.
//
// product gRPC server
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	productpb "calcsvc/gen/grpc/product/pb"
	product "calcsvc/gen/product"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the productpb.ProductServer interface.
type Server struct {
	ProductEndpointH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the product service endpoints.
func New(e *product.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ProductEndpointH: NewProductEndpointHandler(e.ProductEndpoint, uh),
	}
}

// NewProductEndpointHandler creates a gRPC handler which serves the "product"
// service "product" endpoint.
func NewProductEndpointHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeProductEndpointRequest, EncodeProductEndpointResponse)
	}
	return h
}

// ProductEndpoint implements the "ProductEndpoint" method in
// productpb.ProductServer interface.
func (s *Server) ProductEndpoint(ctx context.Context, message *productpb.ProductRequest) (*productpb.ProductResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "product")
	ctx = context.WithValue(ctx, goa.ServiceKey, "product")
	resp, err := s.ProductEndpointH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				er := err.(*product.NotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewProductNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*productpb.ProductResponse), nil
}
