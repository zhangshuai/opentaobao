package wdk

// OrderBuyerInfoBo 结构体
type OrderBuyerInfoBo struct {
	// 收货人名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 收货人电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货人地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// poi经纬度
	Geo string `json:"geo,omitempty" xml:"geo,omitempty"`
	// 配送开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 配送结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}