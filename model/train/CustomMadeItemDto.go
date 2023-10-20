package train

// CustomMadeItemDto 结构体
type CustomMadeItemDto struct {
	// 定制类型
	CustomType string `json:"custom_type,omitempty" xml:"custom_type,omitempty"`
	// 定制车厢集合字符串
	CarriageCustom string `json:"carriage_custom,omitempty" xml:"carriage_custom,omitempty"`
	// 选座坐席号集合字符串
	SeatCustom string `json:"seat_custom,omitempty" xml:"seat_custom,omitempty"`
	// 定制份数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 是否接受非定制,0不接受,1接受
	AcceptNoVipCustom int64 `json:"accept_no_vip_custom,omitempty" xml:"accept_no_vip_custom,omitempty"`
}