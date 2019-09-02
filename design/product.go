
package design

import . "goa.design/goa/v3/dsl"

var _ = Service("product", func() {
	Description("product service.")
	HTTP(func() {
		Path("/api-product")
	})


	Method("product", func() {
		Description("商品信息")
		Payload(func() {
			Field(1, "id", String, "auctionId")
			Required("id")
		})
		Result(Product)
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			GET("/product/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})
})