package trade

// TaobaoKoubeiTribeOpenOrderPageResult 
type TaobaoKoubeiTribeOpenOrderPageResult struct {

    // 订单信息结果
    Data   *OrderInfoResultDto `json:"data,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误提示
    Error   string `json:"error,omitempty"`

    // request唯一ID
    TraceId   string `json:"trace_id,omitempty"`

}