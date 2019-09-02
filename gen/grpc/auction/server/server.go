// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction gRPC server
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	auction "calcsvc/gen/auction"
	auctionpb "calcsvc/gen/grpc/auction/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the auctionpb.AuctionServer interface.
type Server struct {
	PickH goagrpc.UnaryHandler
	GetH  goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the auction service endpoints.
func New(e *auction.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		PickH: NewPickHandler(e.Pick, uh),
		GetH:  NewGetHandler(e.Get, uh),
	}
}

// NewPickHandler creates a gRPC handler which serves the "auction" service
// "pick" endpoint.
func NewPickHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodePickRequest, EncodePickResponse)
	}
	return h
}

// Pick implements the "Pick" method in auctionpb.AuctionServer interface.
func (s *Server) Pick(ctx context.Context, message *auctionpb.PickRequest) (*auctionpb.StoredBottleCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "pick")
	ctx = context.WithValue(ctx, goa.ServiceKey, "auction")
	resp, err := s.PickH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "no_criteria":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "no_match":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*auctionpb.StoredBottleCollection), nil
}

// NewGetHandler creates a gRPC handler which serves the "auction" service
// "get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in auctionpb.AuctionServer interface.
func (s *Server) Get(ctx context.Context, message *auctionpb.GetRequest) (*auctionpb.StoredBottleCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "auction")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*auctionpb.StoredBottleCollection), nil
}