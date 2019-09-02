
package design

import . "goa.design/goa/v3/dsl"

var _ = Service("auction", func() {
	Description("The auction service retrieves bottles given a set of criteria.")
	HTTP(func() {
		Path("/auction")
	})
	Method("pick", func() {
		Payload(Criteria)
		Result(CollectionOf(StoredBottle), func() {
			View("default")
		})
		Error("no_criteria", String, "Missing criteria")
		Error("no_match", String, "No bottle matched given criteria")
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("no_criteria", StatusBadRequest)
			Response("no_match", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("no_criteria", CodeInvalidArgument)
			Response("no_match", CodeNotFound)
		})
	})

	Method("get", func() {
		Description("get")
		Result(CollectionOf(StoredBottle), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/get")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})