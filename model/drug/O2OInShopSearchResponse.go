package drug

// O2oinShopSearchResponse 结构体
type O2oinShopSearchResponse struct {
	// itemList
	ItemList []Itemlist `json:"item_list,omitempty" xml:"item_list>itemlist,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
