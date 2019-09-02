// Code generated by goa v3.0.4, DO NOT EDIT.
//
// product gRPC client types
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	productpb "calcsvc/gen/grpc/product/pb"
	product "calcsvc/gen/product"
	productviews "calcsvc/gen/product/views"
)

// NewProductRequest builds the gRPC request type from the payload of the
// "product" endpoint of the "product" service.
func NewProductRequest(payload *product.ProductPayload) *productpb.ProductRequest {
	message := &productpb.ProductRequest{
		Id: payload.ID,
	}
	return message
}

// NewProductResult builds the result type of the "product" endpoint of the
// "product" service from the gRPC response type.
func NewProductResult(message *productpb.ProductResponse) *productviews.ProductView {
	result := &productviews.ProductView{}
	if message.Id != 0 {
		idptr := int(message.Id)
		result.ID = &idptr
	}
	if message.ArtNo != "" {
		result.ArtNo = &message.ArtNo
	}
	if message.CategoryName != "" {
		result.CategoryName = &message.CategoryName
	}
	if message.CategoryId != 0 {
		categoryIDptr := int(message.CategoryId)
		result.CategoryID = &categoryIDptr
	}
	if message.CertCode != 0 {
		certCodeptr := int(message.CertCode)
		result.CertCode = &certCodeptr
	}
	if message.CertType != 0 {
		certTypeptr := int(message.CertType)
		result.CertType = &certTypeptr
	}
	if message.Colour != "" {
		result.Colour = &message.Colour
	}
	if message.ColourId != 0 {
		colourIDptr := int(message.ColourId)
		result.ColourID = &colourIDptr
	}
	if message.CreatedAt != 0 {
		createdAtptr := int(message.CreatedAt)
		result.CreatedAt = &createdAtptr
	}
	if message.Crowd != "" {
		result.Crowd = &message.Crowd
	}
	if message.CurrentPrice != 0 {
		currentPriceptr := int(message.CurrentPrice)
		result.CurrentPrice = &currentPriceptr
	}
	if message.EndTime != 0 {
		endTimeptr := int(message.EndTime)
		result.EndTime = &endTimeptr
	}
	if message.ExtAttrMap != 0 {
		extAttrMapptr := int(message.ExtAttrMap)
		result.ExtAttrMap = &extAttrMapptr
	}
	if message.Level != 0 {
		levelptr := int(message.Level)
		result.Level = &levelptr
	}
	if message.LevelId != 0 {
		levelIDptr := int(message.LevelId)
		result.LevelID = &levelIDptr
	}
	if message.Name != "" {
		result.Name = &message.Name
	}
	if message.MktPrice != 0 {
		mktPriceptr := int(message.MktPrice)
		result.MktPrice = &mktPriceptr
	}
	if message.OperatorId != 0 {
		operatorIDptr := int(message.OperatorId)
		result.OperatorID = &operatorIDptr
	}
	if message.Recommendations != "" {
		result.Recommendations = &message.Recommendations
	}
	if message.Scenario != "" {
		result.Scenario = &message.Scenario
	}
	if message.Size != "" {
		result.Size = &message.Size
	}
	if message.Status != 0 {
		statusptr := int(message.Status)
		result.Status = &statusptr
	}
	if message.Style != "" {
		result.Style = &message.Style
	}
	if message.StyleId != "" {
		result.StyleID = &message.StyleId
	}
	if message.Summary != "" {
		result.Summary = &message.Summary
	}
	if message.TemplateId != 0 {
		templateIDptr := int(message.TemplateId)
		result.TemplateID = &templateIDptr
	}
	if message.UpdatedAt != 0 {
		updatedAtptr := int(message.UpdatedAt)
		result.UpdatedAt = &updatedAtptr
	}
	if message.Version != 0 {
		versionptr := int(message.Version)
		result.Version = &versionptr
	}
	if message.Weight != "" {
		result.Weight = &message.Weight
	}
	if message.CarouselList != nil {
		result.CarouselList = make([]*productviews.MediaView, len(message.CarouselList))
		for i, val := range message.CarouselList {
			result.CarouselList[i] = &productviews.MediaView{}
			if val.MediaUrl != "" {
				result.CarouselList[i].MediaURL = &val.MediaUrl
			}
			if val.MediaType != 0 {
				mediaTypeptr := int(val.MediaType)
				result.CarouselList[i].MediaType = &mediaTypeptr
			}
			if val.Content != 0 {
				contentptr := int(val.Content)
				result.CarouselList[i].Content = &contentptr
			}
			if val.Sequence != 0 {
				sequenceptr := int(val.Sequence)
				result.CarouselList[i].Sequence = &sequenceptr
			}
			if val.MediaId != 0 {
				mediaIDptr := int(val.MediaId)
				result.CarouselList[i].MediaID = &mediaIDptr
			}
		}
	}
	if message.DetailPics != nil {
		result.DetailPics = make([]*productviews.MediaView, len(message.DetailPics))
		for i, val := range message.DetailPics {
			result.DetailPics[i] = &productviews.MediaView{}
			if val.MediaUrl != "" {
				result.DetailPics[i].MediaURL = &val.MediaUrl
			}
			if val.MediaType != 0 {
				mediaTypeptr := int(val.MediaType)
				result.DetailPics[i].MediaType = &mediaTypeptr
			}
			if val.Content != 0 {
				contentptr := int(val.Content)
				result.DetailPics[i].Content = &contentptr
			}
			if val.Sequence != 0 {
				sequenceptr := int(val.Sequence)
				result.DetailPics[i].Sequence = &sequenceptr
			}
			if val.MediaId != 0 {
				mediaIDptr := int(val.MediaId)
				result.DetailPics[i].MediaID = &mediaIDptr
			}
		}
	}
	if message.ScenarioList != nil {
		result.ScenarioList = make([]string, len(message.ScenarioList))
		for i, val := range message.ScenarioList {
			result.ScenarioList[i] = val
		}
	}
	return result
}

// NewProductNotFoundError builds the error type of the "product" endpoint of
// the "product" service from the gRPC error response type.
func NewProductNotFoundError(message *productpb.ProductNotFoundError) *product.NotFound {
	er := &product.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}