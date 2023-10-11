package ascp

// SignRuleDeleteResultDetail 结构体
type SignRuleDeleteResultDetail struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 收货地：省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 收货地：城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货地：区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收货地：街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}
