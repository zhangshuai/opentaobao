package servicecenter

import (
	"sync"
)

// TmallCarLeaseFreedownpaymentPutResult 结构体
type TmallCarLeaseFreedownpaymentPutResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 当前时间
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// 耗费时间
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// 返回对象,成功：true，失败：false
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallCarLeaseFreedownpaymentPutResult = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseFreedownpaymentPutResult)
	},
}

// GetTmallCarLeaseFreedownpaymentPutResult() 从对象池中获取TmallCarLeaseFreedownpaymentPutResult
func GetTmallCarLeaseFreedownpaymentPutResult() *TmallCarLeaseFreedownpaymentPutResult {
	return poolTmallCarLeaseFreedownpaymentPutResult.Get().(*TmallCarLeaseFreedownpaymentPutResult)
}

// ReleaseTmallCarLeaseFreedownpaymentPutResult 释放TmallCarLeaseFreedownpaymentPutResult
func ReleaseTmallCarLeaseFreedownpaymentPutResult(v *TmallCarLeaseFreedownpaymentPutResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.GmtCurrentTime = 0
	v.CostTime = 0
	v.Object = false
	v.Success = false
	poolTmallCarLeaseFreedownpaymentPutResult.Put(v)
}
