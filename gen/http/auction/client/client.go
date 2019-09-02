// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction client HTTP transport
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

// Client lists the auction service endpoint HTTP clients.
type Client struct {
	// GetAuctionProductListByStatus Doer is the HTTP client used to make requests
	// to the getAuctionProductListByStatus endpoint.
	GetAuctionProductListByStatusDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the auction service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetAuctionProductListByStatusDoer: doer,
		RestoreResponseBody:               restoreBody,
		scheme:                            scheme,
		host:                              host,
		decoder:                           dec,
		encoder:                           enc,
	}
}

// GetAuctionProductListByStatus returns an endpoint that makes HTTP requests
// to the auction service getAuctionProductListByStatus server.
func (c *Client) GetAuctionProductListByStatus() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetAuctionProductListByStatusRequest(c.encoder)
		decodeResponse = DecodeGetAuctionProductListByStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAuctionProductListByStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAuctionProductListByStatusDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("auction", "getAuctionProductListByStatus", err)
		}
		return decodeResponse(resp)
	}
}
