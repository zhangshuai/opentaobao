package tbk

// RightsInstanceCreateResult 结构体
type RightsInstanceCreateResult struct {
	// 淘礼金Id
	Rightsid string `json:"rights_id,omitempty" xml:"rights_id,omitempty"`
	// 淘礼金领取Url
	Sendurl string `json:"send_url,omitempty" xml:"send_url,omitempty"`
	// 投放code【百川商品详情页业务专用】
	Vegascode string `json:"vegas_code,omitempty" xml:"vegas_code,omitempty"`
	// 创建完成后资金账户可用资金，单位元，保留2位小数
	Availablefee string `json:"available_fee,omitempty" xml:"available_fee,omitempty"`
	// 媒体针对此商品今日剩余可领取淘礼金数量
	Itemtodaynumleft int64 `json:"item_today_num_left,omitempty" xml:"item_today_num_left,omitempty"`
}
