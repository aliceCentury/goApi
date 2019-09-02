package design
import . "goa.design/goa/v3/dsl"

var _ = API("cellar", func() {
	Title("Cellar Service")
	Description("HTTP service for managing your wine cellar")

	Server("api", func() {
		Description("cellar hosts the storage, sommelier and swagger services.")
		Services("auction", "user","product", "swagger")
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

var AuctionProduct = ResultType("application/vnd.cellar.stored-bottle", func() {
	Description("拍卖、投标列表.")
	TypeName("AuctionProduct")

	Attributes(func() {
		Attribute("id", String, func() {
			Description("Id")
			Example("50611")
			Meta("rpc:tag", "1")
		})
		Field(2, "add_price", Int,func() {
			Description("加价幅度")
			Example(20)
		})
		Field(3, "art_no",String, func() {
			Description("货号")
			Example("ZS1908038")
		})
		Field(4, "auction_status", Int,func() {
			Description("拍卖结果:  0:成功 1:流拍 2:取当前订单状态   （成功竞得，待用户确认） 3:弃拍")
			Example(1)
		})
		Field(5, "auction_type", Int,func() {
			Description("拍卖类型：0，新手拍(不可加入）；1，直拍(包括0元拍，直接拍,但不可以加入众筹）；2众筹拍（可以加入拍卖众筹）；3，投标拍（特殊的拍卖，可以加入投标众筹）；")
			Example(2)
		})
		Field(6, "bid_scene_id", Int,func() {
			Description("所属投标场Id")
			Example(-1)
		})
		Field(7, "bond_price", Int,func() {
			Description("保证金")
			Example(20)
		})
		Field(8,"buy_number", Int,func() {
			Description("购买数量")
			Example(0)
		})
		Field(9,"buy_unit_price", func() {
			Description("用户购买的单价")
		})
		Field(11,"buyout_price", Int,func() {
			Description("买断价格")
			Example(200)
		})
		Field(12,"cap_price", Int,func() {
			Description("封顶价")
			Example(390)
		})
		Field(13,"crowdfunding_package_id", func() {
			Description("所属众筹包ID")
		})
		Field(14,"current_price", Int,func() {
			Description("当前价")
			Example(110)
		})
		Field(15,"end_time", Int64,func() {
			Example(1567396800000)
		})
		Field(16,"head_portrait")
		Field(17,"is_have_proxy", Int,func() {
			Description("拍卖前有没有代理 0：没有 1：有")
			Example(1)
		})
		Field(18,"is_reserve_price", Int,func() {
			Description("1:有保留价 2:无保留价")
			Example(2)
		})
		Field(19,"last_time", Int64,func() {
			Example(1567396800000)
		})
		Field(20,"limit_number", Int,func() {
			Description("每人限购")
			Example(1)
		})
		Field(21,"mkt_price", Int,func() {
			Description("市场价")
			Example(395)
		})
		Field(22,"pictures_url", String,func() {
			Description("拍卖商品相关图片")
			Example("")
		})
		Field(23,"prod_id", Int32,func() {
			Description("商品ID")
			Example(21320)
		})
		Field(24,"prod_name", String,func() {
			Description("商品名称")
			Example("朱砂平安扣")
		})
		Field(25,"qr_url", String,func() {
			Description("分享二维码")
			Example("")
		})
		Field(26,"remind_time", Int64,func() {
			Description("提醒时间")
			Example(1567389600000)
		})
		Field(27,"reserve_price")
		Field(28,"result_status", Int,func() {
			Description("保留价")
			Example(-1)
		})
		Field(29,"rule_id", Int,func() {
			Description("似乎没用到")
			Example(0)
		})
		Field(30,"serial_num", String,func() {
			Description("拍卖编号")
			Example("10:00~12:00")
		})
		Field(31,"share_url", func() {
			Description("分享图片链接")
		})
		Field(32,"start_auction_price", Int,func() {
			Description("起拍价")
			Example(110)
		})
		Field(33,"start_time", Int64,func() {
			Example(1567389600000)
		})
		Field(34,"title", String,func() {
			Description("拍卖场名称")
			Example("朱砂平安扣")
		})
		Field(35,"total_number", Int,func() {
			Description("总个数")
			Example(1)
		})
		Field(36,"transaction_number", Int,func() {
			Description("成交数量")
			Example(0)
		})
		Field(37,"transaction_price", func() {
			Description("成交总额")
		})
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
	Description("拍卖投标列表的参数")
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

var Product = ResultType("Product", func() {
	Description("商品信息")
	TypeName("Product")

	Attributes(func() {
		Attribute("id", Int, func() {
			Description("商品Id")
			Example(15272)
			Meta("rpc:tag", "1")
		})
		Field(3, "art_no",String,func() {
			Description("编号")
			Example("ZS1908038")
		})
		Field(4, "carouselList", ArrayOf(Media))//"轮播图片（视频）"
		Field(5, "categoryName",String,func() {
			Example("18K金镶嵌翡翠")
		})
		Field(6, "category_id",Int)//"商品分类id"
		Field(7, "cert_code",Int)//证书编号
		Field(8,"cert_type", Int)//证书类型ID
		Field(9,"colour")//颜色
		Field(11,"colour_id",Int)//颜色id
		Field(12,"created_at",Int)
		Field(13,"crowd",String)//属性：适合人群 字段内容格式:1,2,3,男士,女士,学生
		Field(14,"current_price", Int)
		Field(15,"end_time", Int)
		Field(16,"detailPics",ArrayOf(Media))//详情的图片、视频列表
		Field(17,"extAttrMap",Int)
		Field(18,"level", Int)//级别，比如：精品、收藏、经典等
		Field(19,"level_id", Int)
		Field(20,"name", String)//商品名
		Field(21,"mkt_price", Int)
		Field(22,"operator_id", Int)//后台操作者ID
		Field(23,"recommendations",String)
		Field(24,"scenario",String)//后台管理系统
		Field(25,"scenarioList",ArrayOf(String, func() {
			Example("1")
		}))//使用场景(客户端)
		Field(26,"size", String)//大小
		Field(27,"status",Int)//(-1)-已删除 (0)-草稿 (1)-上线（拍品流拍、弃拍或流标、弃标后转为上线） (2)-拍卖中(不一定上线拍卖) (3)-投标中(不一定上线投标) (10)-售出 (11)-下线
		Field(28,"style")
		Field(29,"style_id")//款式ID
		Field(30,"summary",String)//详情介绍
		Field(31,"template_id",Int)//模板id
		Field(32,"updated_at", Int)
		Field(33,"version",Int)
		Field(34,"weight",String)//重量
	})

	//Required("id", "add_price", "title", "vintage")
})

var Media = Type("Media", func() {
	Attribute("media_url", String, "图片URL", func() {
		Example("https://img.cft.upmi.com.cn/o_1d3o3vcmu10723kgta21qb31v25a.jpg")
		Meta("rpc:tag", "5")
	})
	Attribute("media_type", Int, "媒体类型：0-图片，1-视频，2-未知", func() {
		Example(0)
		Meta("rpc:tag", "1")
	})
	Attribute("content", Int, "视频链接", func() {
		Example(1)
		Meta("rpc:tag", "2")
	})
	Attribute("sequence", Int, "上传顺序，默认为0", func() {
		Example(1)
		Meta("rpc:tag", "3")
	})
	Attribute("media_id", Int, "Id", func() {
		Example(169686)
		Meta("rpc:tag", "4")
	})
})