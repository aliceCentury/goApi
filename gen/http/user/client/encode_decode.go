// Code generated by goa v3.0.4, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	"bytes"
	user "calcsvc/gen/user"
	userviews "calcsvc/gen/user/views"
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "user" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the user list
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "list", err)
			}
			p := NewListStoredBottleCollectionOK(body)
			view := "tiny"
			vres := userviews.StoredBottleCollection{p, view}
			if err = userviews.ValidateStoredBottleCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "list", err)
			}
			res := user.NewStoredBottleCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "user" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*user.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "show", "*user.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowUserPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the user show
// server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "show", "*user.ShowPayload", v)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeShowResponse returns a decoder for responses returned by the user show
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeShowResponse may return the following errors:
//	- "not_found" (type *user.NotFound): http.StatusNotFound
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "show", err)
			}
			p := NewShowStoredBottleOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.StoredBottle{p, view}
			if err = userviews.ValidateStoredBottle(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "show", err)
			}
			res := user.NewStoredBottle(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "show", err)
			}
			err = ValidateShowNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "show", err)
			}
			return nil, NewShowNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "user" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the user add server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.Bottle)
		if !ok {
			return goahttp.ErrInvalidType("user", "add", "*user.Bottle", v)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the user add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "user" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*user.RemovePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "remove", "*user.RemovePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveUserPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveResponse returns a decoder for responses returned by the user
// remove endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeRemoveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "remove", resp.StatusCode, string(body))
		}
	}
}

// BuildRateRequest instantiates a HTTP request object with method and path set
// to call the "user" service "rate" endpoint
func (c *Client) BuildRateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RateUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "rate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRateRequest returns an encoder for requests sent to the user rate
// server.
func EncodeRateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(map[uint32][]string)
		if !ok {
			return goahttp.ErrInvalidType("user", "rate", "map[uint32][]string", v)
		}
		values := req.URL.Query()
		for key, value := range p {
			keyStr := strconv.FormatUint(uint64(key), 10)
			for _, val := range value {
				valStr := val
				values.Add(keyStr, valStr)
			}
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeRateResponse returns a decoder for responses returned by the user rate
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeRateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "rate", resp.StatusCode, string(body))
		}
	}
}

// BuildMultiAddRequest instantiates a HTTP request object with method and path
// set to call the "user" service "multi_add" endpoint
func (c *Client) BuildMultiAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiAddUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "multi_add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMultiAddRequest returns an encoder for requests sent to the user
// multi_add server.
func EncodeMultiAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.([]*user.Bottle)
		if !ok {
			return goahttp.ErrInvalidType("user", "multi_add", "[]*user.Bottle", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("user", "multi_add", err)
		}
		return nil
	}
}

// NewUserMultiAddEncoder returns an encoder to encode the multipart request
// for the "user" service "multi_add" endpoint.
func NewUserMultiAddEncoder(encoderFn UserMultiAddEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.([]*user.Bottle)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeMultiAddResponse returns a decoder for responses returned by the user
// multi_add endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeMultiAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "multi_add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "multi_add", resp.StatusCode, string(body))
		}
	}
}

// BuildMultiUpdateRequest instantiates a HTTP request object with method and
// path set to call the "user" service "multi_update" endpoint
func (c *Client) BuildMultiUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiUpdateUserPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "multi_update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMultiUpdateRequest returns an encoder for requests sent to the user
// multi_update server.
func EncodeMultiUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.MultiUpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "multi_update", "*user.MultiUpdatePayload", v)
		}
		values := req.URL.Query()
		for _, value := range p.Ids {
			values.Add("ids", value)
		}
		req.URL.RawQuery = values.Encode()
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("user", "multi_update", err)
		}
		return nil
	}
}

// NewUserMultiUpdateEncoder returns an encoder to encode the multipart request
// for the "user" service "multi_update" endpoint.
func NewUserMultiUpdateEncoder(encoderFn UserMultiUpdateEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.(*user.MultiUpdatePayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeMultiUpdateResponse returns a decoder for responses returned by the
// user multi_update endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeMultiUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "multi_update", resp.StatusCode, string(body))
		}
	}
}

// unmarshalWineryResponseToUserviewsWineryView builds a value of type
// *userviews.WineryView from a value of type *WineryResponse.
func unmarshalWineryResponseToUserviewsWineryView(v *WineryResponse) *userviews.WineryView {
	res := &userviews.WineryView{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// unmarshalComponentResponseToUserviewsComponentView builds a value of type
// *userviews.ComponentView from a value of type *ComponentResponse.
func unmarshalComponentResponseToUserviewsComponentView(v *ComponentResponse) *userviews.ComponentView {
	if v == nil {
		return nil
	}
	res := &userviews.ComponentView{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// unmarshalWineryResponseBodyToUserviewsWineryView builds a value of type
// *userviews.WineryView from a value of type *WineryResponseBody.
func unmarshalWineryResponseBodyToUserviewsWineryView(v *WineryResponseBody) *userviews.WineryView {
	res := &userviews.WineryView{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// unmarshalComponentResponseBodyToUserviewsComponentView builds a value of
// type *userviews.ComponentView from a value of type *ComponentResponseBody.
func unmarshalComponentResponseBodyToUserviewsComponentView(v *ComponentResponseBody) *userviews.ComponentView {
	if v == nil {
		return nil
	}
	res := &userviews.ComponentView{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// marshalUserWineryToWineryRequestBody builds a value of type
// *WineryRequestBody from a value of type *user.Winery.
func marshalUserWineryToWineryRequestBody(v *user.Winery) *WineryRequestBody {
	res := &WineryRequestBody{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// marshalUserComponentToComponentRequestBody builds a value of type
// *ComponentRequestBody from a value of type *user.Component.
func marshalUserComponentToComponentRequestBody(v *user.Component) *ComponentRequestBody {
	if v == nil {
		return nil
	}
	res := &ComponentRequestBody{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// marshalWineryRequestBodyToUserWinery builds a value of type *user.Winery
// from a value of type *WineryRequestBody.
func marshalWineryRequestBodyToUserWinery(v *WineryRequestBody) *user.Winery {
	res := &user.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// marshalComponentRequestBodyToUserComponent builds a value of type
// *user.Component from a value of type *ComponentRequestBody.
func marshalComponentRequestBodyToUserComponent(v *ComponentRequestBody) *user.Component {
	if v == nil {
		return nil
	}
	res := &user.Component{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// marshalUserBottleToBottleRequestBody builds a value of type
// *BottleRequestBody from a value of type *user.Bottle.
func marshalUserBottleToBottleRequestBody(v *user.Bottle) *BottleRequestBody {
	res := &BottleRequestBody{
		Name:        v.Name,
		Vintage:     v.Vintage,
		Description: v.Description,
		Rating:      v.Rating,
	}
	if v.Winery != nil {
		res.Winery = marshalUserWineryToWineryRequestBody(v.Winery)
	}
	if v.Composition != nil {
		res.Composition = make([]*ComponentRequestBody, len(v.Composition))
		for i, val := range v.Composition {
			res.Composition[i] = marshalUserComponentToComponentRequestBody(val)
		}
	}

	return res
}

// marshalBottleRequestBodyToUserBottle builds a value of type *user.Bottle
// from a value of type *BottleRequestBody.
func marshalBottleRequestBodyToUserBottle(v *BottleRequestBody) *user.Bottle {
	res := &user.Bottle{
		Name:        v.Name,
		Vintage:     v.Vintage,
		Description: v.Description,
		Rating:      v.Rating,
	}
	if v.Winery != nil {
		res.Winery = marshalWineryRequestBodyToUserWinery(v.Winery)
	}
	if v.Composition != nil {
		res.Composition = make([]*user.Component, len(v.Composition))
		for i, val := range v.Composition {
			res.Composition[i] = marshalComponentRequestBodyToUserComponent(val)
		}
	}

	return res
}
