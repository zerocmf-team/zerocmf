// Code generated by goctl. DO NOT EDIT.
package types

type CategoryGetReq struct {
	SiteId    int64    `form:"siteId"`
	Name      string   `form:"name,optional"`
	ParentId  int64    `form:"parentId,optional"`
	ListOrder *float64 `form:"listOrder,optional"`
	Status    *int     `form:"status,optional"`
}

type CategoryTreeDataReq struct {
	SiteId   int64  `form:"siteId"`
	Name     string `form:"name,optional"`
	Status   *int   `form:"status,optional"`
	IgnoreId *int64 `form:"ignoreId,optional"`
}

type CategoryShowReq struct {
	SiteId int64 `form:"siteId"`
	Id     int64 `path:"id"`
}

type CategorySaveReq struct {
	SiteId    int64    `form:"siteId"`
	Id        int64    `path:"id,optional"`
	ParentId  int64    `json:"parentId,optional"`
	Name      string   `json:"name"`
	Icon      string   `json:"icon,optional"`
	Desc      string   `json:"desc,optional"`
	ListOrder *float64 `json:"listOrder,optional"`
	Status    *int     `json:"status,optional"`
}

type CategoryDelReq struct {
	SiteId int64 `form:"siteId"`
	Id     int64 `path:"id"`
}

type CategoryBatchDelReq struct {
	SiteId int64 `form:"siteId"`
}

type ProductGetReq struct {
	ProductName     string `form:"productName,optional"`
	ProductCategory int64  `form:"productCategory,optional"`
	Status          *int   `form:"status,optional"`
}

type ProductShowReq struct {
	ProductId int64 `path:"productId"`
}

type ProductSku struct {
	SkuId         int64    `json:"skuId,optional"`
	AttrsVal      string   `json:"attrsVal"`
	SkuCode       string   `json:"skuCode,optional"`
	SkuBarcode    string   `json:"skuBarcode,optional"`
	RetailPrice   float64  `json:"retailPrice"`
	Stock         *float64 `json:"stock,optional"`
	OriginalPrice float64  `json:"originalPrice,optional"`
	CostPrice     float64  `json:"costPrice,optional"`
	Weight        float64  `json:"weight,optional"`
	Status        int      `json:"status,optional"`
}

type AttributesItem struct {
	Name string `json:"name"`
}

type Attributes struct {
	Name  string           `json:"name"`
	Items []AttributesItem `json:"items"`
}

type ProductSaveReq struct {
	ProductId           int64        `path:"productId,optional"`
	ProductName         string       `json:"productName"`
	ProductBarcode      string       `json:"productBarcode,optional"`
	ProductCategory     int64        `json:"productCategory"`
	ProductThumbnail    []string     `json:"productThumbnail"`
	MainVideo           string       `json:"mainVideo,optional"`
	ExplanationVideo    string       `json:"explanationVideo,optional"`
	Attributes          []Attributes `json:"attributes,optional"`
	Price               float64      `json:"price,optional"`
	StockUnit           string       `json:"stockUnit,optional"`
	Stock               int64        `json:"stock,optional"`
	ShareDescription    string       `json:"shareDescription,optional"`
	ProductSellingPoint string       `json:"productSellingPoint,optional"`
	OriginalPrice       float64      `json:"originalPrice,optional"`
	CostPrice           float64      `json:"costPrice,optional"`
	PriceNegotiable     int          `json:"priceNegotiable,optional"`
	HideRemainingStock  int          `json:"hideRemainingStock,optional"`
	ProductContent      string       `json:"productContent,optional"`
	Status              int          `json:"status,optional"`
	ProductSku          []ProductSku `json:"productSku,optional"`
}

type ProductDelReq struct {
}

type ProductBatchDelReq struct {
}
