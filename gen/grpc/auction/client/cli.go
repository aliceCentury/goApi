// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction gRPC client CLI support package
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	auction "calcsvc/gen/auction"
	auctionpb "calcsvc/gen/grpc/auction/pb"
	"encoding/json"
	"fmt"
)

// BuildPickPayload builds the payload for the auction pick endpoint from CLI
// flags.
func BuildPickPayload(auctionPickMessage string) (*auction.Criteria, error) {
	var err error
	var message auctionpb.PickRequest
	{
		if auctionPickMessage != "" {
			err = json.Unmarshal([]byte(auctionPickMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"varietal\": [\n         \"pinot noir\",\n         \"merlot\",\n         \"cabernet franc\"\n      ],\n      \"winery\": \"longoria\"\n   }'")
			}
		}
	}
	v := &auction.Criteria{}
	if message.Name != "" {
		v.Name = &message.Name
	}
	if message.Winery != "" {
		v.Winery = &message.Winery
	}
	if message.Varietal != nil {
		v.Varietal = make([]string, len(message.Varietal))
		for i, val := range message.Varietal {
			v.Varietal[i] = val
		}
	}
	return v, nil
}
