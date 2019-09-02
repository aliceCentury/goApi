// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction endpoints
//
// Command:
// $ goa gen calcsvc/design

package auction

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "auction" service endpoints.
type Endpoints struct {
	Pick goa.Endpoint
	Get  goa.Endpoint
}

// NewEndpoints wraps the methods of the "auction" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Pick: NewPickEndpoint(s),
		Get:  NewGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "auction" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Pick = m(e.Pick)
	e.Get = m(e.Get)
}

// NewPickEndpoint returns an endpoint function that calls the method "pick" of
// service "auction".
func NewPickEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Criteria)
		res, err := s.Pick(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredBottleCollection(res, "default")
		return vres, nil
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "auction".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.Get(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredBottleCollection(res, "default")
		return vres, nil
	}
}
