package lstlogistics2

// MainOrderParam 
type MainOrderParam struct {
    // 主订单id
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    // 发货子订单列表
    SubOrderParamList   []SubOrderParam `json:"sub_order_param_list,omitempty" xml:"sub_order_param_list>sub_order_param,omitempty"`
}