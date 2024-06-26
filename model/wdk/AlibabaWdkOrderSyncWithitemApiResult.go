package wdk

import (
	"sync"
)

// AlibabaWdkOrderSyncWithitemApiResult 结构体
type AlibabaWdkOrderSyncWithitemApiResult struct {
	// 调用接口返回错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口服务内部返回成功，外部success成功只代表接口调用成功，服务内部成功用参数表示
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkOrderSyncWithitemApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderSyncWithitemApiResult)
	},
}

// GetAlibabaWdkOrderSyncWithitemApiResult() 从对象池中获取AlibabaWdkOrderSyncWithitemApiResult
func GetAlibabaWdkOrderSyncWithitemApiResult() *AlibabaWdkOrderSyncWithitemApiResult {
	return poolAlibabaWdkOrderSyncWithitemApiResult.Get().(*AlibabaWdkOrderSyncWithitemApiResult)
}

// ReleaseAlibabaWdkOrderSyncWithitemApiResult 释放AlibabaWdkOrderSyncWithitemApiResult
func ReleaseAlibabaWdkOrderSyncWithitemApiResult(v *AlibabaWdkOrderSyncWithitemApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkOrderSyncWithitemApiResult.Put(v)
}
