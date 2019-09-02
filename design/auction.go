
package design

import . "goa.design/goa/v3/dsl"

var _ = Service("auction", func() {
	Description("The auction service retrieves bottles given a set of criteria.")
	HTTP(func() {
		Path("/api-auction")
	})

	Method("getAuctionProductListByStatus", func() {
		Description("获取拍卖列表")
		Payload(ListData)
		Result(CollectionOf(AuctionProduct), func() {
			View("auctionList")
		})
		HTTP(func() {
			POST("/getAuctionProductListByStatus")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("getAuctionProductDetail", func() {
		Description("拍卖详情")
		Payload(func() {
			Field(1, "id", String, "auctionId")
			Required("id")
		})
		Result(AuctionProduct)
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			GET("/getAuctionProductDetail/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})
})