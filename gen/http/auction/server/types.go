// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction HTTP server types
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	auction "calcsvc/gen/auction"
	auctionviews "calcsvc/gen/auction/views"
)

// GetAuctionProductListByStatusRequestBody is the type of the "auction"
// service "getAuctionProductListByStatus" endpoint HTTP request body.
type GetAuctionProductListByStatusRequestBody struct {
	// 拍卖状态
	AuctionStatus *int `form:"auction_status,omitempty" json:"auction_status,omitempty" xml:"auction_status,omitempty"`
	// 当前页数
	CurrentPage *int `form:"current_page,omitempty" json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页返回的条数
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty" xml:"page_size,omitempty"`
}

// AuctionProductResponseAuctionListCollection is the type of the "auction"
// service "getAuctionProductListByStatus" endpoint HTTP response body.
type AuctionProductResponseAuctionListCollection []*AuctionProductResponseAuctionList

// AuctionProductResponseAuctionList is used to define fields on response body
// types.
type AuctionProductResponseAuctionList struct {
	AddPrice              *int    `form:"add_price,omitempty" json:"add_price,omitempty" xml:"add_price,omitempty"`
	ArtNo                 *string `form:"art_no,omitempty" json:"art_no,omitempty" xml:"art_no,omitempty"`
	AuctionStatus         *int    `form:"auction_status,omitempty" json:"auction_status,omitempty" xml:"auction_status,omitempty"`
	AuctionType           *int    `form:"auction_type,omitempty" json:"auction_type,omitempty" xml:"auction_type,omitempty"`
	BidSceneID            *int    `form:"bid_scene_id,omitempty" json:"bid_scene_id,omitempty" xml:"bid_scene_id,omitempty"`
	BondPrice             *int    `form:"bond_price,omitempty" json:"bond_price,omitempty" xml:"bond_price,omitempty"`
	BuyNumber             *int    `form:"buy_number,omitempty" json:"buy_number,omitempty" xml:"buy_number,omitempty"`
	BuyUnitPrice          *string `form:"buy_unit_price,omitempty" json:"buy_unit_price,omitempty" xml:"buy_unit_price,omitempty"`
	BuyoutPrice           *int    `form:"buyout_price,omitempty" json:"buyout_price,omitempty" xml:"buyout_price,omitempty"`
	CapPrice              *int    `form:"cap_price,omitempty" json:"cap_price,omitempty" xml:"cap_price,omitempty"`
	CrowdfundingPackageID *string `form:"crowdfunding_package_id,omitempty" json:"crowdfunding_package_id,omitempty" xml:"crowdfunding_package_id,omitempty"`
	CurrentPrice          *int    `form:"current_price,omitempty" json:"current_price,omitempty" xml:"current_price,omitempty"`
	EndTime               *int64  `form:"end_time,omitempty" json:"end_time,omitempty" xml:"end_time,omitempty"`
	HeadPortrait          *string `form:"head_portrait,omitempty" json:"head_portrait,omitempty" xml:"head_portrait,omitempty"`
	ID                    *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	IsHaveProxy           *int    `form:"is_have_proxy,omitempty" json:"is_have_proxy,omitempty" xml:"is_have_proxy,omitempty"`
	IsReservePrice        *int    `form:"is_reserve_price,omitempty" json:"is_reserve_price,omitempty" xml:"is_reserve_price,omitempty"`
	LastTime              *int64  `form:"last_time,omitempty" json:"last_time,omitempty" xml:"last_time,omitempty"`
	LimitNumber           *int    `form:"limit_number,omitempty" json:"limit_number,omitempty" xml:"limit_number,omitempty"`
	MktPrice              *int    `form:"mkt_price,omitempty" json:"mkt_price,omitempty" xml:"mkt_price,omitempty"`
	PicturesURL           *string `form:"pictures_url,omitempty" json:"pictures_url,omitempty" xml:"pictures_url,omitempty"`
	ProdID                *int32  `form:"prod_id,omitempty" json:"prod_id,omitempty" xml:"prod_id,omitempty"`
	ProdName              *string `form:"prod_name,omitempty" json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	QrURL                 *string `form:"qr_url,omitempty" json:"qr_url,omitempty" xml:"qr_url,omitempty"`
	RemindTime            *int64  `form:"remind_time,omitempty" json:"remind_time,omitempty" xml:"remind_time,omitempty"`
	ReservePrice          *string `form:"reserve_price,omitempty" json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	ResultStatus          *int    `form:"result_status,omitempty" json:"result_status,omitempty" xml:"result_status,omitempty"`
	SerialNum             *string `form:"serial_num,omitempty" json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	ShareURL              *string `form:"share_url,omitempty" json:"share_url,omitempty" xml:"share_url,omitempty"`
	StartAuctionPrice     *int    `form:"start_auction_price,omitempty" json:"start_auction_price,omitempty" xml:"start_auction_price,omitempty"`
	StartTime             *int64  `form:"start_time,omitempty" json:"start_time,omitempty" xml:"start_time,omitempty"`
	Title                 *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	TotalNumber           *int    `form:"total_number,omitempty" json:"total_number,omitempty" xml:"total_number,omitempty"`
	TransactionNumber     *int    `form:"transaction_number,omitempty" json:"transaction_number,omitempty" xml:"transaction_number,omitempty"`
	TransactionPrice      *string `form:"transaction_price,omitempty" json:"transaction_price,omitempty" xml:"transaction_price,omitempty"`
	UserID                *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName              *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

// NewAuctionProductResponseAuctionListCollection builds the HTTP response body
// from the result of the "getAuctionProductListByStatus" endpoint of the
// "auction" service.
func NewAuctionProductResponseAuctionListCollection(res auctionviews.AuctionProductCollectionView) AuctionProductResponseAuctionListCollection {
	body := make([]*AuctionProductResponseAuctionList, len(res))
	for i, val := range res {
		body[i] = &AuctionProductResponseAuctionList{
			ID:                    val.ID,
			AddPrice:              val.AddPrice,
			ArtNo:                 val.ArtNo,
			AuctionStatus:         val.AuctionStatus,
			AuctionType:           val.AuctionType,
			BidSceneID:            val.BidSceneID,
			BondPrice:             val.BondPrice,
			BuyNumber:             val.BuyNumber,
			BuyUnitPrice:          val.BuyUnitPrice,
			BuyoutPrice:           val.BuyoutPrice,
			CapPrice:              val.CapPrice,
			CrowdfundingPackageID: val.CrowdfundingPackageID,
			CurrentPrice:          val.CurrentPrice,
			EndTime:               val.EndTime,
			HeadPortrait:          val.HeadPortrait,
			IsHaveProxy:           val.IsHaveProxy,
			IsReservePrice:        val.IsReservePrice,
			LastTime:              val.LastTime,
			LimitNumber:           val.LimitNumber,
			MktPrice:              val.MktPrice,
			PicturesURL:           val.PicturesURL,
			ProdID:                val.ProdID,
			ProdName:              val.ProdName,
			QrURL:                 val.QrURL,
			RemindTime:            val.RemindTime,
			ReservePrice:          val.ReservePrice,
			ResultStatus:          val.ResultStatus,
			SerialNum:             val.SerialNum,
			ShareURL:              val.ShareURL,
			StartAuctionPrice:     val.StartAuctionPrice,
			StartTime:             val.StartTime,
			Title:                 val.Title,
			TotalNumber:           val.TotalNumber,
			TransactionNumber:     val.TransactionNumber,
			TransactionPrice:      val.TransactionPrice,
			UserID:                val.UserID,
			UserName:              val.UserName,
		}
	}
	return body
}

// NewGetAuctionProductListByStatusListData builds a auction service
// getAuctionProductListByStatus endpoint payload.
func NewGetAuctionProductListByStatusListData(body *GetAuctionProductListByStatusRequestBody) *auction.ListData {
	v := &auction.ListData{
		AuctionStatus: body.AuctionStatus,
		CurrentPage:   body.CurrentPage,
		PageSize:      body.PageSize,
	}
	return v
}
