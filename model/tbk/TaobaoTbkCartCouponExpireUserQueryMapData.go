package tbk

// TaobaotbkcartcouponexpireuserqueryMapData 结构体
type TaobaotbkcartcouponexpireuserqueryMapData struct {
	// 商品ID对应的sku集合
	Skuidlist []int64 `json:"sku_id_list,omitempty" xml:"sku_id_list>int64,omitempty"`
	// 商品ID
	Itemid int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
