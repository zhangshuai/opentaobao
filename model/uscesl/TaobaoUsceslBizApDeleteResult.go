package uscesl

import (
	"sync"
)

// TaobaoUsceslBizApDeleteResult 结构体
type TaobaoUsceslBizApDeleteResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务执行返回码
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 返回码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回执行对象
	Target bool `json:"target,omitempty" xml:"target,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoUsceslBizApDeleteResult = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizApDeleteResult)
	},
}

// GetTaobaoUsceslBizApDeleteResult() 从对象池中获取TaobaoUsceslBizApDeleteResult
func GetTaobaoUsceslBizApDeleteResult() *TaobaoUsceslBizApDeleteResult {
	return poolTaobaoUsceslBizApDeleteResult.Get().(*TaobaoUsceslBizApDeleteResult)
}

// ReleaseTaobaoUsceslBizApDeleteResult 释放TaobaoUsceslBizApDeleteResult
func ReleaseTaobaoUsceslBizApDeleteResult(v *TaobaoUsceslBizApDeleteResult) {
	v.Message = ""
	v.BusinessCode = ""
	v.ReturnCode = 0
	v.Target = false
	v.Success = false
	poolTaobaoUsceslBizApDeleteResult.Put(v)
}
