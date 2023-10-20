package servicecenter

import (
	"sync"
)

// TpFundsSendQuery 结构体
type TpFundsSendQuery struct {
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 资金权益类型。1：预付款红包；2：尾款红包
	FundsType int64 `json:"funds_type,omitempty" xml:"funds_type,omitempty"`
}

var poolTpFundsSendQuery = sync.Pool{
	New: func() any {
		return new(TpFundsSendQuery)
	},
}

// GetTpFundsSendQuery() 从对象池中获取TpFundsSendQuery
func GetTpFundsSendQuery() *TpFundsSendQuery {
	return poolTpFundsSendQuery.Get().(*TpFundsSendQuery)
}

// ReleaseTpFundsSendQuery 释放TpFundsSendQuery
func ReleaseTpFundsSendQuery(v *TpFundsSendQuery) {
	v.BizOrderId = 0
	v.FundsType = 0
	poolTpFundsSendQuery.Put(v)
}
