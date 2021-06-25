package bill

// BillDto 
type BillDto struct {

    // 账单金额,退款时返回的是负数
    TotalAmount   int64 `json:"total_amount,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 0-未支付,1-支付成功,2-支付失败
    Status   int64 `json:"status,omitempty"`

    // 交易子订单编号
    OrderId   string `json:"order_id,omitempty"`

    // 业务时间,订单交易完成的时间
    BizTime   string `json:"biz_time,omitempty"`

    // 账单金额,退款时返回的是负数
    Amount   int64 `json:"amount,omitempty"`

    // 支付宝备注
    AlipayNotice   string `json:"alipay_notice,omitempty"`

    // 支付宝商户订单号
    AlipayOutno   string `json:"alipay_outno,omitempty"`

    // 目标支付宝账户编号
    ObjAlipayId   string `json:"obj_alipay_id,omitempty"`

    // 交易订单编号
    TradeId   string `json:"trade_id,omitempty"`

    // 支付宝账户编号
    AlipayId   string `json:"alipay_id,omitempty"`

    // 科目编号
    AccountId   int64 `json:"account_id,omitempty"`

    // 支付宝账户名称
    AlipayMail   string `json:"alipay_mail,omitempty"`

    // 目标支付宝账户名称
    ObjAlipayMail   string `json:"obj_alipay_mail,omitempty"`

    // 支付时间,收取佣金时支付宝的扣款时间
    PayTime   string `json:"pay_time,omitempty"`

    // 账单编号
    Bid   int64 `json:"bid,omitempty"`

    // 记账时间
    BookTime   string `json:"book_time,omitempty"`

    // 所属商品编号
    NumIid   string `json:"num_iid,omitempty"`

    // 支付宝交易号,暂不提供
    AlipayNo   string `json:"alipay_no,omitempty"`

}