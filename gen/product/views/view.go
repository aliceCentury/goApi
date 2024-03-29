// Code generated by goa v3.0.4, DO NOT EDIT.
//
// product views
//
// Command:
// $ goa gen calcsvc/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Product is the viewed result type that is projected based on a view.
type Product struct {
	// Type to project
	Projected *ProductView
	// View to render
	View string
}

// ProductView is a type that runs validations on a projected type.
type ProductView struct {
	// 商品Id
	ID *int
	// 编号
	ArtNo           *string
	CarouselList    []*MediaView
	CategoryName    *string
	CategoryID      *int
	CertCode        *int
	CertType        *int
	Colour          *string
	ColourID        *int
	CreatedAt       *int
	Crowd           *string
	CurrentPrice    *int
	EndTime         *int
	DetailPics      []*MediaView
	ExtAttrMap      *int
	Level           *int
	LevelID         *int
	Name            *string
	MktPrice        *int
	OperatorID      *int
	Recommendations *string
	Scenario        *string
	ScenarioList    []string
	Size            *string
	Status          *int
	Style           *string
	StyleID         *string
	Summary         *string
	TemplateID      *int
	UpdatedAt       *int
	Version         *int
	Weight          *string
}

// MediaView is a type that runs validations on a projected type.
type MediaView struct {
	// 图片URL
	MediaURL *string
	// 媒体类型：0-图片，1-视频，2-未知
	MediaType *int
	// 视频链接
	Content *int
	// 上传顺序，默认为0
	Sequence *int
	// Id
	MediaID *int
}

var (
	// ProductMap is a map of attribute names in result type Product indexed by
	// view name.
	ProductMap = map[string][]string{
		"default": []string{
			"id",
			"art_no",
			"carouselList",
			"categoryName",
			"category_id",
			"cert_code",
			"cert_type",
			"colour",
			"colour_id",
			"created_at",
			"crowd",
			"current_price",
			"end_time",
			"detailPics",
			"extAttrMap",
			"level",
			"level_id",
			"name",
			"mkt_price",
			"operator_id",
			"recommendations",
			"scenario",
			"scenarioList",
			"size",
			"status",
			"style",
			"style_id",
			"summary",
			"template_id",
			"updated_at",
			"version",
			"weight",
		},
	}
)

// ValidateProduct runs the validations defined on the viewed result type
// Product.
func ValidateProduct(result *Product) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProductView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateProductView runs the validations defined on ProductView using the
// "default" view.
func ValidateProductView(result *ProductView) (err error) {

	return
}

// ValidateMediaView runs the validations defined on MediaView.
func ValidateMediaView(result *MediaView) (err error) {

	return
}
