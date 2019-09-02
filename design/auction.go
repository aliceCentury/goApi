
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
})