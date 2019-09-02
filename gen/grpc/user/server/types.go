// Code generated by goa v3.0.4, DO NOT EDIT.
//
// user gRPC server types
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	userpb "calcsvc/gen/grpc/user/pb"
	user "calcsvc/gen/user"
	userviews "calcsvc/gen/user/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewStoredBottleCollection builds the gRPC response type from the result of
// the "list" endpoint of the "user" service.
func NewStoredBottleCollection(result userviews.StoredBottleCollectionView) *userpb.StoredBottleCollection {
	message := &userpb.StoredBottleCollection{}
	message.Field = make([]*userpb.StoredBottle, len(result))
	for i, val := range result {
		message.Field[i] = &userpb.StoredBottle{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.Name != nil {
			message.Field[i].Name = *val.Name
		}
		if val.Vintage != nil {
			message.Field[i].Vintage = *val.Vintage
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.Rating != nil {
			message.Field[i].Rating = *val.Rating
		}
		if val.Winery != nil {
			message.Field[i].Winery = svcUserviewsWineryViewToUserpbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*userpb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &userpb.Component{}
				if val.Varietal != nil {
					message.Field[i].Composition[j].Varietal = *val.Varietal
				}
				if val.Percentage != nil {
					message.Field[i].Composition[j].Percentage = *val.Percentage
				}
			}
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the "user"
// service from the gRPC request type.
func NewShowPayload(message *userpb.ShowRequest, view *string) *user.ShowPayload {
	v := &user.ShowPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewShowResponse builds the gRPC response type from the result of the "show"
// endpoint of the "user" service.
func NewShowResponse(result *userviews.StoredBottleView) *userpb.ShowResponse {
	message := &userpb.ShowResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.Name != nil {
		message.Name = *result.Name
	}
	if result.Vintage != nil {
		message.Vintage = *result.Vintage
	}
	if result.Description != nil {
		message.Description = *result.Description
	}
	if result.Rating != nil {
		message.Rating = *result.Rating
	}
	if result.Winery != nil {
		message.Winery = svcUserviewsWineryViewToUserpbWinery(result.Winery)
	}
	if result.Composition != nil {
		message.Composition = make([]*userpb.Component, len(result.Composition))
		for i, val := range result.Composition {
			message.Composition[i] = &userpb.Component{}
			if val.Varietal != nil {
				message.Composition[i].Varietal = *val.Varietal
			}
			if val.Percentage != nil {
				message.Composition[i].Percentage = *val.Percentage
			}
		}
	}
	return message
}

// NewShowNotFoundError builds the gRPC error response type from the error of
// the "show" endpoint of the "user" service.
func NewShowNotFoundError(er *user.NotFound) *userpb.ShowNotFoundError {
	message := &userpb.ShowNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "user" service
// from the gRPC request type.
func NewAddPayload(message *userpb.AddRequest) *user.Bottle {
	v := &user.Bottle{
		Name:    message.Name,
		Vintage: message.Vintage,
	}
	if message.Description != "" {
		v.Description = &message.Description
	}
	if message.Rating != 0 {
		v.Rating = &message.Rating
	}
	if message.Winery != nil {
		v.Winery = protobufUserpbWineryToUserWinery(message.Winery)
	}
	if message.Composition != nil {
		v.Composition = make([]*user.Component, len(message.Composition))
		for i, val := range message.Composition {
			v.Composition[i] = &user.Component{
				Varietal: val.Varietal,
			}
			if val.Percentage != 0 {
				v.Composition[i].Percentage = &val.Percentage
			}
		}
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "user" service.
func NewAddResponse(result string) *userpb.AddResponse {
	message := &userpb.AddResponse{}
	message.Field = result
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the "user"
// service from the gRPC request type.
func NewRemovePayload(message *userpb.RemoveRequest) *user.RemovePayload {
	v := &user.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "user" service.
func NewRemoveResponse() *userpb.RemoveResponse {
	message := &userpb.RemoveResponse{}
	return message
}

// NewRatePayload builds the payload of the "rate" endpoint of the "user"
// service from the gRPC request type.
func NewRatePayload(message *userpb.RateRequest) map[uint32][]string {
	v := make(map[uint32][]string, len(message.Field))
	for key, val := range message.Field {
		tk := key
		tv := make([]string, len(val.Field))
		for i, val := range val.Field {
			tv[i] = val
		}
		v[tk] = tv
	}
	return v
}

// NewRateResponse builds the gRPC response type from the result of the "rate"
// endpoint of the "user" service.
func NewRateResponse() *userpb.RateResponse {
	message := &userpb.RateResponse{}
	return message
}

// NewMultiAddPayload builds the payload of the "multi_add" endpoint of the
// "user" service from the gRPC request type.
func NewMultiAddPayload(message *userpb.MultiAddRequest) []*user.Bottle {
	v := make([]*user.Bottle, len(message.Field))
	for i, val := range message.Field {
		v[i] = &user.Bottle{
			Name:    val.Name,
			Vintage: val.Vintage,
		}
		if val.Description != "" {
			v[i].Description = &val.Description
		}
		if val.Rating != 0 {
			v[i].Rating = &val.Rating
		}
		if val.Winery != nil {
			v[i].Winery = protobufUserpbWineryToUserWinery(val.Winery)
		}
		if val.Composition != nil {
			v[i].Composition = make([]*user.Component, len(val.Composition))
			for j, val := range val.Composition {
				v[i].Composition[j] = &user.Component{
					Varietal: val.Varietal,
				}
				if val.Percentage != 0 {
					v[i].Composition[j].Percentage = &val.Percentage
				}
			}
		}
	}
	return v
}

// NewMultiAddResponse builds the gRPC response type from the result of the
// "multi_add" endpoint of the "user" service.
func NewMultiAddResponse(result []string) *userpb.MultiAddResponse {
	message := &userpb.MultiAddResponse{}
	message.Field = make([]string, len(result))
	for i, val := range result {
		message.Field[i] = val
	}
	return message
}

// NewMultiUpdatePayload builds the payload of the "multi_update" endpoint of
// the "user" service from the gRPC request type.
func NewMultiUpdatePayload(message *userpb.MultiUpdateRequest) *user.MultiUpdatePayload {
	v := &user.MultiUpdatePayload{}
	if message.Ids != nil {
		v.Ids = make([]string, len(message.Ids))
		for i, val := range message.Ids {
			v.Ids[i] = val
		}
	}
	if message.Bottles != nil {
		v.Bottles = make([]*user.Bottle, len(message.Bottles))
		for i, val := range message.Bottles {
			v.Bottles[i] = &user.Bottle{
				Name:    val.Name,
				Vintage: val.Vintage,
			}
			if val.Description != "" {
				v.Bottles[i].Description = &val.Description
			}
			if val.Rating != 0 {
				v.Bottles[i].Rating = &val.Rating
			}
			if val.Winery != nil {
				v.Bottles[i].Winery = protobufUserpbWineryToUserWinery(val.Winery)
			}
			if val.Composition != nil {
				v.Bottles[i].Composition = make([]*user.Component, len(val.Composition))
				for j, val := range val.Composition {
					v.Bottles[i].Composition[j] = &user.Component{
						Varietal: val.Varietal,
					}
					if val.Percentage != 0 {
						v.Bottles[i].Composition[j].Percentage = &val.Percentage
					}
				}
			}
		}
	}
	return v
}

// NewMultiUpdateResponse builds the gRPC response type from the result of the
// "multi_update" endpoint of the "user" service.
func NewMultiUpdateResponse() *userpb.MultiUpdateResponse {
	message := &userpb.MultiUpdateResponse{}
	return message
}

// ValidateWinery runs the validations defined on Winery.
func ValidateWinery(message *userpb.Winery) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.region", message.Region, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.country", message.Country, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateComponent runs the validations defined on Component.
func ValidateComponent(message *userpb.Component) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.varietal", message.Varietal, "[A-Za-z' ]+"))
	if utf8.RuneCountInString(message.Varietal) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.varietal", message.Varietal, utf8.RuneCountInString(message.Varietal), 100, false))
	}
	if message.Percentage < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.percentage", message.Percentage, 1, true))
	}
	if message.Percentage > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.percentage", message.Percentage, 100, false))
	}
	return
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *userpb.AddRequest) (err error) {
	if message.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "message"))
	}
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if message.Description != "" {
		if utf8.RuneCountInString(message.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 2000, false))
		}
	}
	if message.Rating != 0 {
		if message.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 1, true))
		}
	}
	if message.Rating != 0 {
		if message.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 5, false))
		}
	}
	return
}

// ValidateRateRequest runs the validations defined on RateRequest.
func ValidateRateRequest(message *userpb.RateRequest) (err error) {
	if message.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("field", "message"))
	}
	for _, v := range message.Field {
		if v != nil {
			if err2 := ValidateArrayOfString(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateArrayOfString runs the validations defined on ArrayOfString.
func ValidateArrayOfString(message *userpb.ArrayOfString) (err error) {
	if message.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("field", "message"))
	}
	return
}

// ValidateMultiAddRequest runs the validations defined on MultiAddRequest.
func ValidateMultiAddRequest(message *userpb.MultiAddRequest) (err error) {
	if message.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("field", "message"))
	}
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateBottle runs the validations defined on Bottle.
func ValidateBottle(message *userpb.Bottle) (err error) {
	if message.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "message"))
	}
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if utf8.RuneCountInString(message.Description) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 2000, false))
	}
	if message.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 1, true))
	}
	if message.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 5, false))
	}
	return
}

// ValidateMultiUpdateRequest runs the validations defined on
// MultiUpdateRequest.
func ValidateMultiUpdateRequest(message *userpb.MultiUpdateRequest) (err error) {
	if message.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ids", "message"))
	}
	if message.Bottles == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bottles", "message"))
	}
	for _, e := range message.Bottles {
		if e != nil {
			if err2 := ValidateBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// svcUserviewsWineryViewToUserpbWinery builds a value of type *userpb.Winery
// from a value of type *userviews.WineryView.
func svcUserviewsWineryViewToUserpbWinery(v *userviews.WineryView) *userpb.Winery {
	res := &userpb.Winery{}
	if v.Name != nil {
		res.Name = *v.Name
	}
	if v.Region != nil {
		res.Region = *v.Region
	}
	if v.Country != nil {
		res.Country = *v.Country
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}

// protobufUserpbWineryToUserviewsWineryView builds a value of type
// *userviews.WineryView from a value of type *userpb.Winery.
func protobufUserpbWineryToUserviewsWineryView(v *userpb.Winery) *userviews.WineryView {
	res := &userviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}

// protobufUserpbWineryToUserWinery builds a value of type *user.Winery from a
// value of type *userpb.Winery.
func protobufUserpbWineryToUserWinery(v *userpb.Winery) *user.Winery {
	res := &user.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}

// svcUserWineryToUserpbWinery builds a value of type *userpb.Winery from a
// value of type *user.Winery.
func svcUserWineryToUserpbWinery(v *user.Winery) *userpb.Winery {
	res := &userpb.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}