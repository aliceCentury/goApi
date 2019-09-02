// Code generated by goa v3.0.4, DO NOT EDIT.
//
// auction views
//
// Command:
// $ goa gen calcsvc/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// AuctionProductCollection is the viewed result type that is projected based
// on a view.
type AuctionProductCollection struct {
	// Type to project
	Projected AuctionProductCollectionView
	// View to render
	View string
}

// AuctionProduct is the viewed result type that is projected based on a view.
type AuctionProduct struct {
	// Type to project
	Projected *AuctionProductView
	// View to render
	View string
}

// AuctionProductCollectionView is a type that runs validations on a projected
// type.
type AuctionProductCollectionView []*AuctionProductView

// AuctionProductView is a type that runs validations on a projected type.
type AuctionProductView struct {
	// Id
	ID *string
	// 加价幅度
	AddPrice *int
	// 货号
	ArtNo *string
	// 拍卖结果:  0:成功 1:流拍 2:取当前订单状态   （成功竞得，待用户确认） 3:弃拍
	AuctionStatus *int
	// 拍卖类型：0，新手拍(不可加入）；1，直拍(包括0元拍，直接拍,但不可以加入众筹）；2众筹拍（可以加入拍卖众筹）；3，投标拍（特殊的拍卖，可以加入投标众筹）；
	AuctionType *int
	// 所属投标场Id
	BidSceneID *int
	// 保证金
	BondPrice *int
	// 购买数量
	BuyNumber *int
	// 用户购买的单价
	BuyUnitPrice *string
	// 买断价格
	BuyoutPrice *int
	// 封顶价
	CapPrice *int
	// 所属众筹包ID
	CrowdfundingPackageID *string
	// 当前价
	CurrentPrice *int
	EndTime      *int64
	HeadPortrait *string
	// 拍卖前有没有代理 0：没有 1：有
	IsHaveProxy *int
	// 1:有保留价 2:无保留价
	IsReservePrice *int
	LastTime       *int64
	// 每人限购
	LimitNumber *int
	// 市场价
	MktPrice *int
	// 拍卖商品相关图片
	PicturesURL *string
	// 商品ID
	ProdID *int32
	// 商品名称
	ProdName *string
	// 分享二维码
	QrURL *string
	// 提醒时间
	RemindTime   *int64
	ReservePrice *string
	// 保留价
	ResultStatus *int
	// 似乎没用到
	RuleID *int
	// 拍卖编号
	SerialNum *string
	// 分享图片链接
	ShareURL *string
	// 起拍价
	StartAuctionPrice *int
	StartTime         *int64
	// 拍卖场名称
	Title *string
	// 总个数
	TotalNumber *int
	// 成交数量
	TransactionNumber *int
	// 成交总额
	TransactionPrice *string
	UserID           *string
	UserName         *string
}

var (
	// AuctionProductCollectionMap is a map of attribute names in result type
	// AuctionProductCollection indexed by view name.
	AuctionProductCollectionMap = map[string][]string{
		"bid": []string{
			"add_price",
			"art_no",
			"head_portrait",
		},
		"auctionList": []string{
			"add_price",
			"art_no",
			"auction_status",
			"auction_type",
			"bid_scene_id",
			"bond_price",
			"buy_number",
			"buy_unit_price",
			"buyout_price",
			"cap_price",
			"crowdfunding_package_id",
			"current_price",
			"end_time",
			"head_portrait",
			"id",
			"is_have_proxy",
			"is_reserve_price",
			"last_time",
			"limit_number",
			"mkt_price",
			"pictures_url",
			"prod_id",
			"prod_name",
			"qr_url",
			"remind_time",
			"reserve_price",
			"result_status",
			"serial_num",
			"share_url",
			"start_auction_price",
			"start_time",
			"title",
			"total_number",
			"transaction_number",
			"transaction_price",
			"user_id",
			"user_name",
		},
		"default": []string{
			"id",
			"add_price",
			"art_no",
			"auction_status",
			"auction_type",
			"bid_scene_id",
			"bond_price",
			"buy_number",
			"buy_unit_price",
			"buyout_price",
			"cap_price",
			"crowdfunding_package_id",
			"current_price",
			"end_time",
			"head_portrait",
			"is_have_proxy",
			"is_reserve_price",
			"last_time",
			"limit_number",
			"mkt_price",
			"pictures_url",
			"prod_id",
			"prod_name",
			"qr_url",
			"remind_time",
			"reserve_price",
			"result_status",
			"rule_id",
			"serial_num",
			"share_url",
			"start_auction_price",
			"start_time",
			"title",
			"total_number",
			"transaction_number",
			"transaction_price",
			"user_id",
			"user_name",
		},
	}
	// AuctionProductMap is a map of attribute names in result type AuctionProduct
	// indexed by view name.
	AuctionProductMap = map[string][]string{
		"bid": []string{
			"add_price",
			"art_no",
			"head_portrait",
		},
		"auctionList": []string{
			"add_price",
			"art_no",
			"auction_status",
			"auction_type",
			"bid_scene_id",
			"bond_price",
			"buy_number",
			"buy_unit_price",
			"buyout_price",
			"cap_price",
			"crowdfunding_package_id",
			"current_price",
			"end_time",
			"head_portrait",
			"id",
			"is_have_proxy",
			"is_reserve_price",
			"last_time",
			"limit_number",
			"mkt_price",
			"pictures_url",
			"prod_id",
			"prod_name",
			"qr_url",
			"remind_time",
			"reserve_price",
			"result_status",
			"serial_num",
			"share_url",
			"start_auction_price",
			"start_time",
			"title",
			"total_number",
			"transaction_number",
			"transaction_price",
			"user_id",
			"user_name",
		},
		"default": []string{
			"id",
			"add_price",
			"art_no",
			"auction_status",
			"auction_type",
			"bid_scene_id",
			"bond_price",
			"buy_number",
			"buy_unit_price",
			"buyout_price",
			"cap_price",
			"crowdfunding_package_id",
			"current_price",
			"end_time",
			"head_portrait",
			"is_have_proxy",
			"is_reserve_price",
			"last_time",
			"limit_number",
			"mkt_price",
			"pictures_url",
			"prod_id",
			"prod_name",
			"qr_url",
			"remind_time",
			"reserve_price",
			"result_status",
			"rule_id",
			"serial_num",
			"share_url",
			"start_auction_price",
			"start_time",
			"title",
			"total_number",
			"transaction_number",
			"transaction_price",
			"user_id",
			"user_name",
		},
	}
)

// ValidateAuctionProductCollection runs the validations defined on the viewed
// result type AuctionProductCollection.
func ValidateAuctionProductCollection(result AuctionProductCollection) (err error) {
	switch result.View {
	case "bid":
		err = ValidateAuctionProductCollectionViewBid(result.Projected)
	case "auctionList":
		err = ValidateAuctionProductCollectionViewAuctionList(result.Projected)
	case "default", "":
		err = ValidateAuctionProductCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"bid", "auctionList", "default"})
	}
	return
}

// ValidateAuctionProduct runs the validations defined on the viewed result
// type AuctionProduct.
func ValidateAuctionProduct(result *AuctionProduct) (err error) {
	switch result.View {
	case "bid":
		err = ValidateAuctionProductViewBid(result.Projected)
	case "auctionList":
		err = ValidateAuctionProductViewAuctionList(result.Projected)
	case "default", "":
		err = ValidateAuctionProductView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"bid", "auctionList", "default"})
	}
	return
}

// ValidateAuctionProductCollectionViewBid runs the validations defined on
// AuctionProductCollectionView using the "bid" view.
func ValidateAuctionProductCollectionViewBid(result AuctionProductCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateAuctionProductViewBid(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateAuctionProductCollectionViewAuctionList runs the validations defined
// on AuctionProductCollectionView using the "auctionList" view.
func ValidateAuctionProductCollectionViewAuctionList(result AuctionProductCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateAuctionProductViewAuctionList(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateAuctionProductCollectionView runs the validations defined on
// AuctionProductCollectionView using the "default" view.
func ValidateAuctionProductCollectionView(result AuctionProductCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateAuctionProductView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateAuctionProductViewBid runs the validations defined on
// AuctionProductView using the "bid" view.
func ValidateAuctionProductViewBid(result *AuctionProductView) (err error) {

	return
}

// ValidateAuctionProductViewAuctionList runs the validations defined on
// AuctionProductView using the "auctionList" view.
func ValidateAuctionProductViewAuctionList(result *AuctionProductView) (err error) {

	return
}

// ValidateAuctionProductView runs the validations defined on
// AuctionProductView using the "default" view.
func ValidateAuctionProductView(result *AuctionProductView) (err error) {

	return
}
