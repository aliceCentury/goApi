package design
import . "goa.design/goa/v3/dsl"

var _ = API("cellar", func() {
	Title("Cellar Service")
	Description("HTTP service for managing your wine cellar")

	Server("api", func() {
		Description("cellar hosts the storage, sommelier and swagger services.")
		Services("auction", "user", "swagger")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000/api")
			URI("grpc://localhost:8080/api")
		})
		Host("goa.design", func() {
			Description("public host")
			URI("https://goa.design/cellar")
			URI("grpcs://goa.design/cellar")
		})
	})
})

var StoredBottle = ResultType("StoredBottle", func() {
	Description("A StoredBottle describes a bottle retrieved by the storage service.")
	Reference(Bottle)
	TypeName("StoredBottle")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the bottle.", func() {
			Example("123abc")
			Meta("rpc:tag", "8")
		})
		Field(2, "name")
		Field(3, "winery")
		Field(4, "vintage")
		Field(5, "composition")
		Field(6, "description")
		Field(7, "rating")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("winery", func() {
			View("tiny")
		})
		Attribute("vintage")
		Attribute("composition")
		Attribute("description")
		Attribute("rating")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("winery", func() {
			View("tiny")
		})
	})

	Required("id", "name", "winery", "vintage")
})

var Bottle = Type("Bottle", func() {
	Description("Bottle describes a bottle of wine to be stored.")
	Attribute("name", String, "Name of bottle", func() {
		MaxLength(100)
		Example("Blue's Cuvee")
		Meta("rpc:tag", "1")
	})
	Field(2, "winery", Winery, "Winery that produces wine")
	Attribute("vintage", UInt32, "Vintage of bottle", func() {
		Minimum(1900)
		Maximum(2020)
		Meta("rpc:tag", "3")
	})
	Field(4, "composition", ArrayOf(Component), "Composition is the list of grape varietals and associated percentage.")
	Attribute("description", String, "Description of bottle", func() {
		MaxLength(2000)
		Example("Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah")
		Meta("rpc:tag", "5")
	})
	Attribute("rating", UInt32, "Rating of bottle from 1 (worst) to 5 (best)", func() {
		Minimum(1)
		Maximum(5)
		Meta("rpc:tag", "6")
	})
	Required("name", "winery", "vintage")
})

var Winery = ResultType("Winery", func() {
	Attributes(func() {
		Attribute("name", String, "Name of winery", func() {
			Example("Longoria")
			Meta("rpc:tag", "1")
		})
		Attribute("region", String, "Region of winery", func() {
			Pattern(`(?i)[a-z '\.]+`)
			Example("Central Coast, California")
			Meta("rpc:tag", "2")
		})
		Attribute("country", String, "Country of winery", func() {
			Pattern(`(?i)[a-z '\.]+`)
			Example("USA")
			Meta("rpc:tag", "3")
		})
		Attribute("url", String, "Winery website URL", func() {
			Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
			Example("http://www.longoriawine.com/")
			Meta("rpc:tag", "4")
		})
	})
	View("default", func() {
		Attribute("name")
		Attribute("region")
		Attribute("country")
		Attribute("url")
	})
	View("tiny", func() {
		Attribute("name")
	})
	Required("name", "region", "country")
})

var Component = Type("Component", func() {
	Attribute("varietal", String, "Grape varietal", func() {
		Pattern(`[A-Za-z' ]+`)
		MaxLength(100)
		Example("Syrah")
		Meta("rpc:tag", "1")
	})
	Attribute("percentage", UInt32, "Percentage of varietal in wine", func() {
		Minimum(1)
		Maximum(100)
		Meta("rpc:tag", "2")
	})
	Required("varietal")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a bottle that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("bottle 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing bottle")
	Required("message", "id")
})

var Criteria = Type("Criteria", func() {
	Description("Criteria described a set of criteria used to pick a bottle. All criteria are optional, at least one must be provided.")
	Attribute("name", String, "Name of bottle to pick", func() {
		Example("Blue's Cuvee")
		Meta("rpc:tag", "1")
	})
	Attribute("varietal", ArrayOf(String), "Varietals in preference order", func() {
		Example([]string{"pinot noir", "merlot", "cabernet franc"})
		Meta("rpc:tag", "2")
	})
	Attribute("winery", String, "Winery of bottle to pick", func() {
		Example("longoria")
		Meta("rpc:tag", "3")
	})
})


var AuctionProduct = ResultType("application/vnd.cellar.stored-bottle", func() {
	Description("A StoredBottle describes a bottle retrieved by the storage service.")
	Reference(Bottle)
	TypeName("AuctionProduct")

	Attributes(func() {
		Attribute("id", String, func() {
			Example("50611")
			Meta("rpc:tag", "8")
		})
		Field(2, "add_price", Int,func() {
			Example(20)
		})
		Field(3, "art_no",String, func() {
			Example("ZS1908038")
		})
		Field(4, "auction_status", Int,func() {
			Example(1)
		})
		Field(5, "auction_type", Int,func() {
			Example(2)
		})
		Field(6, "bid_scene_id", Int,func() {
			Example(-1)
		})
		Field(7, "bond_price", Int,func() {
			Example(20)
		})
		Field(8,"buy_number", Int,func() {
			Example(0)
		})
		Field(9,"buy_unit_price")
		Field(10,"buy_number", Int,func() {
			Example(0)
		})
		Field(11,"buyout_price", Int,func() {
			Example(200)
		})
		Field(12,"cap_price", Int,func() {
			Example(390)
		})
		Field(13,"crowdfunding_package_id")
		Field(14,"current_price", Int,func() {
			Example(110)
		})
		Field(15,"end_time", Int64,func() {
			Example(1567396800000)
		})
		Field(16,"head_portrait")
		Field(17,"is_have_proxy", Int,func() {
			Example(1)
		})
		Field(18,"is_reserve_price", Int,func() {
			Example(2)
		})
		Field(19,"last_time", Int64,func() {
			Example(1567396800000)
		})
		Field(20,"limit_number", Int,func() {
			Example(1)
		})
		Field(21,"mkt_price", Int,func() {
			Example(395)
		})
		Field(22,"pictures_url", String,func() {
			Example("")
		})
		Field(23,"prod_id", Int32,func() {
			Example(21320)
		})
		Field(24,"prod_name", String,func() {
			Example("朱砂平安扣")
		})
		Field(25,"qr_url", String,func() {
			Example("")
		})
		Field(26,"remind_time", Int64,func() {
			Example(1567389600000)
		})
		Field(27,"reserve_price")
		Field(28,"result_status", Int,func() {
			Example(-1)
		})
		Field(29,"rule_id", Int,func() {
			Example(0)
		})
		Field(30,"serial_num", String,func() {
			Example("10:00~12:00")
		})
		Field(31,"share_url")
		Field(32,"start_auction_price", Int,func() {
			Example(110)
		})
		Field(33,"start_time", Int64,func() {
			Example(1567389600000)
		})
		Field(34,"title", String,func() {
			Example("朱砂平安扣")
		})
		Field(35,"total_number", Int,func() {
			Example(1)
		})
		Field(36,"transaction_number", Int,func() {
			Example(0)
		})
		Field(37,"transaction_price")
		Field(38,"user_id")
		Field(39,"user_name")
	})

	View("bid", func() {
		Attribute("add_price")
		Attribute("art_no")
		Attribute("head_portrait", func() {
			View("auctionList")
		})
	})

	View("auctionList", func() {
		Attribute("add_price")
		Attribute("art_no")
		Attribute("auction_status")
		Attribute("auction_type")
		Attribute("bid_scene_id")
		Attribute("bond_price")
		Attribute("buy_number")
		Attribute("buy_unit_price")
		Attribute("buyout_price")
		Attribute("cap_price")
		Attribute("crowdfunding_package_id")
		Attribute("current_price")
		Attribute("end_time")
		Attribute("head_portrait")
		Attribute("id")
		Attribute("is_have_proxy")
		Attribute("is_reserve_price")
		Attribute("last_time")
		Attribute("limit_number")
		Attribute("mkt_price")
		Attribute("pictures_url")
		Attribute("prod_id")
		Attribute("prod_name")
		Attribute("qr_url")
		Attribute("remind_time")
		Attribute("reserve_price")
		Attribute("result_status")
		Attribute("serial_num")
		Attribute("share_url")
		Attribute("start_auction_price")
		Attribute("start_time")
		Attribute("title")
		Attribute("total_number")
		Attribute("transaction_number")
		Attribute("transaction_price")
		Attribute("user_id")
		Attribute("user_name")
	})

	//Required("id", "add_price", "title", "vintage")
})
var ListData = Type("ListData", func() {
	Description("拍卖投标列表的参数")//{"auction_status":2,"current_page":1,"page_size":30}
	Attribute("auction_status", Int, "拍卖状态 1:历史 2:正在进行 3:即将开始 ", func() {
		Example(2)
		Meta("rpc:tag", "1")
	})
	Attribute("current_page", Int, "当前页数", func() {
		Example(1)
		Meta("rpc:tag", "2")
	})
	Attribute("page_size", Int, "每页返回的条数", func() {
		Example(30)
		Meta("rpc:tag", "3")
	})
})