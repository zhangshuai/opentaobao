package tbk

// OrderPage 结构体
type OrderPage struct {
	// PublisherOrderDto
	Results []PublisherOrderDto `json:"results,omitempty" xml:"results>publisher_order_dto,omitempty"`
	// 真正的业务数据结构
	Result []PublisherRefundOrderDto `json:"result,omitempty" xml:"result>publisher_refund_order_dto,omitempty"`
	// 位点字段，由调用方原样传递
	Positionindex string `json:"position_index,omitempty" xml:"position_index,omitempty"`
	// 页码
	Pageno int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	Pagesize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 上一页
	Prepage int64 `json:"pre_page,omitempty" xml:"pre_page,omitempty"`
	// 下一页
	Nextpage int64 `json:"next_page,omitempty" xml:"next_page,omitempty"`
	// 是否还有上一页
	Haspre bool `json:"has_pre,omitempty" xml:"has_pre,omitempty"`
	// 是否还有下一页
	Hasnext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
