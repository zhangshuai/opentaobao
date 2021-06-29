package ascpchannel

// ExternalCreateRefundOrderDetailRequest 
type ExternalCreateRefundOrderDetailRequest struct {
    // 币种
    CurrencyType   string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
    // 子销售单号
    SubSaleOrderNo   string `json:"sub_sale_order_no,omitempty" xml:"sub_sale_order_no,omitempty"`
    // 退款金额
    RefundFee   int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 退款类型
    RefundType   int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    // 销售订单号
    SaleOrderNo   string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
    // 外部退款单号
    OutRefundNo   string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 退货数量
    RefundQuantity   int64 `json:"refund_quantity,omitempty" xml:"refund_quantity,omitempty"`
    // 外部子订单号
    OutSubOrderNo   string `json:"out_sub_order_no,omitempty" xml:"out_sub_order_no,omitempty"`
    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
}
