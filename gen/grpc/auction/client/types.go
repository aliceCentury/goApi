// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction gRPC client types
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	auction "calcsvc/gen/auction"
	auctionviews "calcsvc/gen/auction/views"
	auctionpb "calcsvc/gen/grpc/auction/pb"
)

// NewGetAuctionProductListByStatusRequest builds the gRPC request type from
// the payload of the "getAuctionProductListByStatus" endpoint of the "auction"
// service.
func NewGetAuctionProductListByStatusRequest(payload *auction.ListData) *auctionpb.GetAuctionProductListByStatusRequest {
	message := &auctionpb.GetAuctionProductListByStatusRequest{}
	if payload.AuctionStatus != nil {
		message.AuctionStatus = int32(*payload.AuctionStatus)
	}
	if payload.CurrentPage != nil {
		message.CurrentPage = int32(*payload.CurrentPage)
	}
	if payload.PageSize != nil {
		message.PageSize = int32(*payload.PageSize)
	}
	return message
}

// NewGetAuctionProductListByStatusResult builds the result type of the
// "getAuctionProductListByStatus" endpoint of the "auction" service from the
// gRPC response type.
func NewGetAuctionProductListByStatusResult(message *auctionpb.AuctionProductCollection) auctionviews.AuctionProductCollectionView {
	result := make([]*auctionviews.AuctionProductView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &auctionviews.AuctionProductView{}
		if val.Id != "" {
			result[i].ID = &val.Id
		}
		if val.AddPrice != 0 {
			addPriceptr := int(val.AddPrice)
			result[i].AddPrice = &addPriceptr
		}
		if val.ArtNo != "" {
			result[i].ArtNo = &val.ArtNo
		}
		if val.AuctionStatus != 0 {
			auctionStatusptr := int(val.AuctionStatus)
			result[i].AuctionStatus = &auctionStatusptr
		}
		if val.AuctionType != 0 {
			auctionTypeptr := int(val.AuctionType)
			result[i].AuctionType = &auctionTypeptr
		}
		if val.BidSceneId != 0 {
			bidSceneIDptr := int(val.BidSceneId)
			result[i].BidSceneID = &bidSceneIDptr
		}
		if val.BondPrice != 0 {
			bondPriceptr := int(val.BondPrice)
			result[i].BondPrice = &bondPriceptr
		}
		if val.BuyNumber != 0 {
			buyNumberptr := int(val.BuyNumber)
			result[i].BuyNumber = &buyNumberptr
		}
		if val.BuyUnitPrice != "" {
			result[i].BuyUnitPrice = &val.BuyUnitPrice
		}
		if val.BuyoutPrice != 0 {
			buyoutPriceptr := int(val.BuyoutPrice)
			result[i].BuyoutPrice = &buyoutPriceptr
		}
		if val.CapPrice != 0 {
			capPriceptr := int(val.CapPrice)
			result[i].CapPrice = &capPriceptr
		}
		if val.CrowdfundingPackageId != "" {
			result[i].CrowdfundingPackageID = &val.CrowdfundingPackageId
		}
		if val.CurrentPrice != 0 {
			currentPriceptr := int(val.CurrentPrice)
			result[i].CurrentPrice = &currentPriceptr
		}
		if val.EndTime != 0 {
			result[i].EndTime = &val.EndTime
		}
		if val.HeadPortrait != "" {
			result[i].HeadPortrait = &val.HeadPortrait
		}
		if val.IsHaveProxy != 0 {
			isHaveProxyptr := int(val.IsHaveProxy)
			result[i].IsHaveProxy = &isHaveProxyptr
		}
		if val.IsReservePrice != 0 {
			isReservePriceptr := int(val.IsReservePrice)
			result[i].IsReservePrice = &isReservePriceptr
		}
		if val.LastTime != 0 {
			result[i].LastTime = &val.LastTime
		}
		if val.LimitNumber != 0 {
			limitNumberptr := int(val.LimitNumber)
			result[i].LimitNumber = &limitNumberptr
		}
		if val.MktPrice != 0 {
			mktPriceptr := int(val.MktPrice)
			result[i].MktPrice = &mktPriceptr
		}
		if val.PicturesUrl != "" {
			result[i].PicturesURL = &val.PicturesUrl
		}
		if val.ProdId != 0 {
			result[i].ProdID = &val.ProdId
		}
		if val.ProdName != "" {
			result[i].ProdName = &val.ProdName
		}
		if val.QrUrl != "" {
			result[i].QrURL = &val.QrUrl
		}
		if val.RemindTime != 0 {
			result[i].RemindTime = &val.RemindTime
		}
		if val.ReservePrice != "" {
			result[i].ReservePrice = &val.ReservePrice
		}
		if val.ResultStatus != 0 {
			resultStatusptr := int(val.ResultStatus)
			result[i].ResultStatus = &resultStatusptr
		}
		if val.RuleId != 0 {
			ruleIDptr := int(val.RuleId)
			result[i].RuleID = &ruleIDptr
		}
		if val.SerialNum != "" {
			result[i].SerialNum = &val.SerialNum
		}
		if val.ShareUrl != "" {
			result[i].ShareURL = &val.ShareUrl
		}
		if val.StartAuctionPrice != 0 {
			startAuctionPriceptr := int(val.StartAuctionPrice)
			result[i].StartAuctionPrice = &startAuctionPriceptr
		}
		if val.StartTime != 0 {
			result[i].StartTime = &val.StartTime
		}
		if val.Title != "" {
			result[i].Title = &val.Title
		}
		if val.TotalNumber != 0 {
			totalNumberptr := int(val.TotalNumber)
			result[i].TotalNumber = &totalNumberptr
		}
		if val.TransactionNumber != 0 {
			transactionNumberptr := int(val.TransactionNumber)
			result[i].TransactionNumber = &transactionNumberptr
		}
		if val.TransactionPrice != "" {
			result[i].TransactionPrice = &val.TransactionPrice
		}
		if val.UserId != "" {
			result[i].UserID = &val.UserId
		}
		if val.UserName != "" {
			result[i].UserName = &val.UserName
		}
	}
	return result
}

// NewGetAuctionProductDetailRequest builds the gRPC request type from the
// payload of the "getAuctionProductDetail" endpoint of the "auction" service.
func NewGetAuctionProductDetailRequest(payload *auction.GetAuctionProductDetailPayload) *auctionpb.GetAuctionProductDetailRequest {
	message := &auctionpb.GetAuctionProductDetailRequest{
		Id: payload.ID,
	}
	return message
}

// NewGetAuctionProductDetailResult builds the result type of the
// "getAuctionProductDetail" endpoint of the "auction" service from the gRPC
// response type.
func NewGetAuctionProductDetailResult(message *auctionpb.GetAuctionProductDetailResponse) *auctionviews.AuctionProductView {
	result := &auctionviews.AuctionProductView{}
	if message.Id != "" {
		result.ID = &message.Id
	}
	if message.AddPrice != 0 {
		addPriceptr := int(message.AddPrice)
		result.AddPrice = &addPriceptr
	}
	if message.ArtNo != "" {
		result.ArtNo = &message.ArtNo
	}
	if message.AuctionStatus != 0 {
		auctionStatusptr := int(message.AuctionStatus)
		result.AuctionStatus = &auctionStatusptr
	}
	if message.AuctionType != 0 {
		auctionTypeptr := int(message.AuctionType)
		result.AuctionType = &auctionTypeptr
	}
	if message.BidSceneId != 0 {
		bidSceneIDptr := int(message.BidSceneId)
		result.BidSceneID = &bidSceneIDptr
	}
	if message.BondPrice != 0 {
		bondPriceptr := int(message.BondPrice)
		result.BondPrice = &bondPriceptr
	}
	if message.BuyNumber != 0 {
		buyNumberptr := int(message.BuyNumber)
		result.BuyNumber = &buyNumberptr
	}
	if message.BuyUnitPrice != "" {
		result.BuyUnitPrice = &message.BuyUnitPrice
	}
	if message.BuyoutPrice != 0 {
		buyoutPriceptr := int(message.BuyoutPrice)
		result.BuyoutPrice = &buyoutPriceptr
	}
	if message.CapPrice != 0 {
		capPriceptr := int(message.CapPrice)
		result.CapPrice = &capPriceptr
	}
	if message.CrowdfundingPackageId != "" {
		result.CrowdfundingPackageID = &message.CrowdfundingPackageId
	}
	if message.CurrentPrice != 0 {
		currentPriceptr := int(message.CurrentPrice)
		result.CurrentPrice = &currentPriceptr
	}
	if message.EndTime != 0 {
		result.EndTime = &message.EndTime
	}
	if message.HeadPortrait != "" {
		result.HeadPortrait = &message.HeadPortrait
	}
	if message.IsHaveProxy != 0 {
		isHaveProxyptr := int(message.IsHaveProxy)
		result.IsHaveProxy = &isHaveProxyptr
	}
	if message.IsReservePrice != 0 {
		isReservePriceptr := int(message.IsReservePrice)
		result.IsReservePrice = &isReservePriceptr
	}
	if message.LastTime != 0 {
		result.LastTime = &message.LastTime
	}
	if message.LimitNumber != 0 {
		limitNumberptr := int(message.LimitNumber)
		result.LimitNumber = &limitNumberptr
	}
	if message.MktPrice != 0 {
		mktPriceptr := int(message.MktPrice)
		result.MktPrice = &mktPriceptr
	}
	if message.PicturesUrl != "" {
		result.PicturesURL = &message.PicturesUrl
	}
	if message.ProdId != 0 {
		result.ProdID = &message.ProdId
	}
	if message.ProdName != "" {
		result.ProdName = &message.ProdName
	}
	if message.QrUrl != "" {
		result.QrURL = &message.QrUrl
	}
	if message.RemindTime != 0 {
		result.RemindTime = &message.RemindTime
	}
	if message.ReservePrice != "" {
		result.ReservePrice = &message.ReservePrice
	}
	if message.ResultStatus != 0 {
		resultStatusptr := int(message.ResultStatus)
		result.ResultStatus = &resultStatusptr
	}
	if message.RuleId != 0 {
		ruleIDptr := int(message.RuleId)
		result.RuleID = &ruleIDptr
	}
	if message.SerialNum != "" {
		result.SerialNum = &message.SerialNum
	}
	if message.ShareUrl != "" {
		result.ShareURL = &message.ShareUrl
	}
	if message.StartAuctionPrice != 0 {
		startAuctionPriceptr := int(message.StartAuctionPrice)
		result.StartAuctionPrice = &startAuctionPriceptr
	}
	if message.StartTime != 0 {
		result.StartTime = &message.StartTime
	}
	if message.Title != "" {
		result.Title = &message.Title
	}
	if message.TotalNumber != 0 {
		totalNumberptr := int(message.TotalNumber)
		result.TotalNumber = &totalNumberptr
	}
	if message.TransactionNumber != 0 {
		transactionNumberptr := int(message.TransactionNumber)
		result.TransactionNumber = &transactionNumberptr
	}
	if message.TransactionPrice != "" {
		result.TransactionPrice = &message.TransactionPrice
	}
	if message.UserId != "" {
		result.UserID = &message.UserId
	}
	if message.UserName != "" {
		result.UserName = &message.UserName
	}
	return result
}

// NewGetAuctionProductDetailNotFoundError builds the error type of the
// "getAuctionProductDetail" endpoint of the "auction" service from the gRPC
// error response type.
func NewGetAuctionProductDetailNotFoundError(message *auctionpb.GetAuctionProductDetailNotFoundError) *auction.NotFound {
	er := &auction.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}
