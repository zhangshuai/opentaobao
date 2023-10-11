package ascp

// UnbundleItemmappingRequest 结构体
type UnbundleItemmappingRequest struct {
	// 业务请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 货主编码；优仓分配，长度不会超过32位，不含特殊字符
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 业务请求时间(毫秒数)
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
