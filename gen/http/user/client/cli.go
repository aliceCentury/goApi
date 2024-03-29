// Code generated by goa v3.0.4, DO NOT EDIT.
//
// user HTTP client CLI support package
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	user "calcsvc/gen/user"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildShowPayload builds the payload for the user show endpoint from CLI
// flags.
func BuildShowPayload(userShowID string, userShowView string) (*user.ShowPayload, error) {
	var id string
	{
		id = userShowID
	}
	var view *string
	{
		if userShowView != "" {
			view = &userShowView
		}
	}
	payload := &user.ShowPayload{
		ID:   id,
		View: view,
	}
	return payload, nil
}

// BuildAddPayload builds the payload for the user add endpoint from CLI flags.
func BuildAddPayload(userAddBody string) (*user.Bottle, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(userAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"composition\": [\n         {\n            \"percentage\": 64,\n            \"varietal\": \"Syrah\"\n         },\n         {\n            \"percentage\": 64,\n            \"varietal\": \"Syrah\"\n         }\n      ],\n      \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n      \"name\": \"Blue\\'s Cuvee\",\n      \"rating\": 1,\n      \"vintage\": 1998,\n      \"winery\": {\n         \"country\": \"USA\",\n         \"name\": \"Longoria\",\n         \"region\": \"Central Coast, California\",\n         \"url\": \"http://www.longoriawine.com/\"\n      }\n   }'")
		}
		if body.Winery == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
		}
		if utf8.RuneCountInString(body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
		}
		if body.Winery != nil {
			if err2 := ValidateWineryRequestBody(body.Winery); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", body.Vintage, 1900, true))
		}
		if body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", body.Vintage, 2020, false))
		}
		for _, e := range body.Composition {
			if e != nil {
				if err2 := ValidateComponentRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if body.Description != nil {
			if utf8.RuneCountInString(*body.Description) > 2000 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
			}
		}
		if body.Rating != nil {
			if *body.Rating < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
			}
		}
		if body.Rating != nil {
			if *body.Rating > 5 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &user.Bottle{
		Name:        body.Name,
		Vintage:     body.Vintage,
		Description: body.Description,
		Rating:      body.Rating,
	}
	if body.Winery != nil {
		v.Winery = marshalWineryRequestBodyToUserWinery(body.Winery)
	}
	if body.Composition != nil {
		v.Composition = make([]*user.Component, len(body.Composition))
		for i, val := range body.Composition {
			v.Composition[i] = marshalComponentRequestBodyToUserComponent(val)
		}
	}
	return v, nil
}

// BuildRemovePayload builds the payload for the user remove endpoint from CLI
// flags.
func BuildRemovePayload(userRemoveID string) (*user.RemovePayload, error) {
	var id string
	{
		id = userRemoveID
	}
	payload := &user.RemovePayload{
		ID: id,
	}
	return payload, nil
}

// BuildMultiAddPayload builds the payload for the user multi_add endpoint from
// CLI flags.
func BuildMultiAddPayload(userMultiAddBody string) ([]*user.Bottle, error) {
	var err error
	var body []*BottleRequestBody
	{
		err = json.Unmarshal([]byte(userMultiAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'[\n      {\n         \"composition\": [\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            },\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            }\n         ],\n         \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"rating\": 2,\n         \"vintage\": 1945,\n         \"winery\": {\n            \"country\": \"USA\",\n            \"name\": \"Longoria\",\n            \"region\": \"Central Coast, California\",\n            \"url\": \"http://www.longoriawine.com/\"\n         }\n      },\n      {\n         \"composition\": [\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            },\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            }\n         ],\n         \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"rating\": 2,\n         \"vintage\": 1945,\n         \"winery\": {\n            \"country\": \"USA\",\n            \"name\": \"Longoria\",\n            \"region\": \"Central Coast, California\",\n            \"url\": \"http://www.longoriawine.com/\"\n         }\n      },\n      {\n         \"composition\": [\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            },\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            }\n         ],\n         \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"rating\": 2,\n         \"vintage\": 1945,\n         \"winery\": {\n            \"country\": \"USA\",\n            \"name\": \"Longoria\",\n            \"region\": \"Central Coast, California\",\n            \"url\": \"http://www.longoriawine.com/\"\n         }\n      },\n      {\n         \"composition\": [\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            },\n            {\n               \"percentage\": 64,\n               \"varietal\": \"Syrah\"\n            }\n         ],\n         \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n         \"name\": \"Blue\\'s Cuvee\",\n         \"rating\": 2,\n         \"vintage\": 1945,\n         \"winery\": {\n            \"country\": \"USA\",\n            \"name\": \"Longoria\",\n            \"region\": \"Central Coast, California\",\n            \"url\": \"http://www.longoriawine.com/\"\n         }\n      }\n   ]'")
		}
	}
	v := make([]*user.Bottle, len(body))
	for i, val := range body {
		v[i] = &user.Bottle{
			Name:        val.Name,
			Vintage:     val.Vintage,
			Description: val.Description,
			Rating:      val.Rating,
		}
		if val.Winery != nil {
			v[i].Winery = marshalWineryRequestBodyToUserWinery(val.Winery)
		}
		if val.Composition != nil {
			v[i].Composition = make([]*user.Component, len(val.Composition))
			for j, val := range val.Composition {
				v[i].Composition[j] = marshalComponentRequestBodyToUserComponent(val)
			}
		}
	}
	return v, nil
}

// BuildMultiUpdatePayload builds the payload for the user multi_update
// endpoint from CLI flags.
func BuildMultiUpdatePayload(userMultiUpdateBody string, userMultiUpdateIds string) (*user.MultiUpdatePayload, error) {
	var err error
	var body MultiUpdateRequestBody
	{
		err = json.Unmarshal([]byte(userMultiUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"bottles\": [\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 64,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 64,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 1945,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         },\n         {\n            \"composition\": [\n               {\n                  \"percentage\": 64,\n                  \"varietal\": \"Syrah\"\n               },\n               {\n                  \"percentage\": 64,\n                  \"varietal\": \"Syrah\"\n               }\n            ],\n            \"description\": \"Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah\",\n            \"name\": \"Blue\\'s Cuvee\",\n            \"rating\": 2,\n            \"vintage\": 1945,\n            \"winery\": {\n               \"country\": \"USA\",\n               \"name\": \"Longoria\",\n               \"region\": \"Central Coast, California\",\n               \"url\": \"http://www.longoriawine.com/\"\n            }\n         }\n      ]\n   }'")
		}
		if body.Bottles == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("bottles", "body"))
		}
		for _, e := range body.Bottles {
			if e != nil {
				if err2 := ValidateBottleRequestBody(e); err2 != nil {
					err = goa.MergeErrors(err, err2)
				}
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var ids []string
	{
		err = json.Unmarshal([]byte(userMultiUpdateIds), &ids)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for ids, example of valid JSON:\n%s", "'[\n      \"Tenetur consectetur aut deserunt.\",\n      \"Non quae atque animi qui.\"\n   ]'")
		}
	}
	v := &user.MultiUpdatePayload{}
	if body.Bottles != nil {
		v.Bottles = make([]*user.Bottle, len(body.Bottles))
		for i, val := range body.Bottles {
			v.Bottles[i] = marshalBottleRequestBodyToUserBottle(val)
		}
	}
	v.Ids = ids
	return v, nil
}
