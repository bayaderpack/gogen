package models

type TypeValue string

const (
	P TypeValue = "P"
	F TypeValue = "F"
)

type ProductEnum string

const (
	ProductCategory    ProductEnum = "product_category"
	ProductTag         ProductEnum = "product_tag"
	ProductFilter      ProductEnum = "product_filter"
	ProductBrand       ProductEnum = "product_brand"
	ProductSuitability ProductEnum = "product_suitability"
	ProductAttribute   ProductEnum = "product_attribute"
	MediaTag           ProductEnum = "media_tag"
	MediaCategory      ProductEnum = "media_category"
)
