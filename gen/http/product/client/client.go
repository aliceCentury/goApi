// Code generated by goa v3.0.4, DO NOT EDIT.
//
// product client HTTP transport
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the product service endpoint HTTP clients.
type Client struct {
	// ProductEndpoint Doer is the HTTP client used to make requests to the product
	// endpoint.
	ProductEndpointDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the product service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ProductEndpointDoer: doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// ProductEndpoint returns an endpoint that makes HTTP requests to the product
// service product server.
func (c *Client) ProductEndpoint() goa.Endpoint {
	var (
		decodeResponse = DecodeProductEndpointResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildProductEndpointRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ProductEndpointDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("product", "product", err)
		}
		return decodeResponse(resp)
	}
}
