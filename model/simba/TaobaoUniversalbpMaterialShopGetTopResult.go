package simba

// TaobaouniversalbpmaterialshopgetTopResult 结构体
type TaobaouniversalbpmaterialshopgetTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	ShopInfoVO *ShopInfoVo `json:"shop_info_v_o,omitempty" xml:"shop_info_v_o,omitempty"`
}
