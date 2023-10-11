package wlbimports

// OrderHandoverContentSynopsisDto 结构体
type OrderHandoverContentSynopsisDto struct {
	// 预约大包类型：TRAY：托，SACK：麻袋（非自寄模式必填）
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 预约大包类型名称：托、麻袋（非自寄模式必填）
	ContentTypeName string `json:"content_type_name,omitempty" xml:"content_type_name,omitempty"`
	// 数量（非自寄模式必填）
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
