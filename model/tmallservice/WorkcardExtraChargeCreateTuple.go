package tmallservice

// WorkcardExtraChargeCreateTuple 结构体
type WorkcardExtraChargeCreateTuple struct {
	// 图片地址回传集合
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 费用项名称
	ChargeItemName string `json:"charge_item_name,omitempty" xml:"charge_item_name,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 费用项单价(分)
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
}
