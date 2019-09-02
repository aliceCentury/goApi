// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction gRPC server encoders and decoders
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	auction "calcsvc/gen/auction"
	auctionviews "calcsvc/gen/auction/views"
	auctionpb "calcsvc/gen/grpc/auction/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeGetAuctionProductListByStatusResponse encodes responses from the
// "auction" service "getAuctionProductListByStatus" endpoint.
func EncodeGetAuctionProductListByStatusResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(auctionviews.AuctionProductCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("auction", "getAuctionProductListByStatus", "auctionviews.AuctionProductCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewAuctionProductCollection(result)
	return resp, nil
}

// DecodeGetAuctionProductListByStatusRequest decodes requests sent to
// "auction" service "getAuctionProductListByStatus" endpoint.
func DecodeGetAuctionProductListByStatusRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *auctionpb.GetAuctionProductListByStatusRequest
		ok      bool
	)
	{
		if message, ok = v.(*auctionpb.GetAuctionProductListByStatusRequest); !ok {
			return nil, goagrpc.ErrInvalidType("auction", "getAuctionProductListByStatus", "*auctionpb.GetAuctionProductListByStatusRequest", v)
		}
	}
	var payload *auction.ListData
	{
		payload = NewGetAuctionProductListByStatusPayload(message)
	}
	return payload, nil
}
