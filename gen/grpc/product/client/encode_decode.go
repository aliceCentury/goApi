// Code generated by goa v3.0.4, DO NOT EDIT.
//
// product gRPC client encoders and decoders
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	productpb "calcsvc/gen/grpc/product/pb"
	product "calcsvc/gen/product"
	productviews "calcsvc/gen/product/views"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildProductEndpointFunc builds the remote method to invoke for "product"
// service "product" endpoint.
func BuildProductEndpointFunc(grpccli productpb.ProductClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.ProductEndpoint(ctx, reqpb.(*productpb.ProductRequest), opts...)
		}
		return grpccli.ProductEndpoint(ctx, &productpb.ProductRequest{}, opts...)
	}
}

// EncodeProductEndpointRequest encodes requests sent to product product
// endpoint.
func EncodeProductEndpointRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*product.ProductPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("product", "product", "*product.ProductPayload", v)
	}
	return NewProductRequest(payload), nil
}

// DecodeProductEndpointResponse decodes responses from the product product
// endpoint.
func DecodeProductEndpointResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*productpb.ProductResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("product", "product", "*productpb.ProductResponse", v)
	}
	res := NewProductResult(message)
	vres := &productviews.Product{Projected: res, View: view}
	if err := productviews.ValidateProduct(vres); err != nil {
		return nil, err
	}
	return product.NewProduct(vres), nil
}
