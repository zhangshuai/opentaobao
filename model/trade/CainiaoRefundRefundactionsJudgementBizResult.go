package trade

import (
	"sync"
)

// CainiaoRefundRefundactionsJudgementBizResult 结构体
type CainiaoRefundRefundactionsJudgementBizResult struct {
	// 调用时错误码
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 调用时错误描述
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 是否能对某一个订单进行退货退款操作的返回值
	Data *OrderRefundOperationResponse `json:"data,omitempty" xml:"data,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoRefundRefundactionsJudgementBizResult = sync.Pool{
	New: func() any {
		return new(CainiaoRefundRefundactionsJudgementBizResult)
	},
}

// GetCainiaoRefundRefundactionsJudgementBizResult() 从对象池中获取CainiaoRefundRefundactionsJudgementBizResult
func GetCainiaoRefundRefundactionsJudgementBizResult() *CainiaoRefundRefundactionsJudgementBizResult {
	return poolCainiaoRefundRefundactionsJudgementBizResult.Get().(*CainiaoRefundRefundactionsJudgementBizResult)
}

// ReleaseCainiaoRefundRefundactionsJudgementBizResult 释放CainiaoRefundRefundactionsJudgementBizResult
func ReleaseCainiaoRefundRefundactionsJudgementBizResult(v *CainiaoRefundRefundactionsJudgementBizResult) {
	v.StatusCode = ""
	v.StatusMessage = ""
	v.Data = nil
	v.Success = false
	poolCainiaoRefundRefundactionsJudgementBizResult.Put(v)
}
