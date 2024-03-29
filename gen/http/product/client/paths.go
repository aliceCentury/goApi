// Code generated by goa v3.0.4, DO NOT EDIT.
//
// HTTP request path constructors for the product service.
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	"fmt"
)

// ProductEndpointProductPath returns the URL path to the product service product HTTP endpoint.
func ProductEndpointProductPath(id string) string {
	return fmt.Sprintf("/api-product/product/%v", id)
}
