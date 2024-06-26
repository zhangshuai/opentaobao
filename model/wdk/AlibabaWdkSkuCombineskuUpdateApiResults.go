package wdk

import (
	"sync"
)

// AlibabaWdkSkuCombineskuUpdateApiResults 结构体
type AlibabaWdkSkuCombineskuUpdateApiResults struct {
	// 商品列表
	Models []AlibabaWdkSkuCombineskuUpdateApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_combinesku_update_api_result,omitempty"`
	// 接口调用异常编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 接口调用异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuCombineskuUpdateApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCombineskuUpdateApiResults)
	},
}

// GetAlibabaWdkSkuCombineskuUpdateApiResults() 从对象池中获取AlibabaWdkSkuCombineskuUpdateApiResults
func GetAlibabaWdkSkuCombineskuUpdateApiResults() *AlibabaWdkSkuCombineskuUpdateApiResults {
	return poolAlibabaWdkSkuCombineskuUpdateApiResults.Get().(*AlibabaWdkSkuCombineskuUpdateApiResults)
}

// ReleaseAlibabaWdkSkuCombineskuUpdateApiResults 释放AlibabaWdkSkuCombineskuUpdateApiResults
func ReleaseAlibabaWdkSkuCombineskuUpdateApiResults(v *AlibabaWdkSkuCombineskuUpdateApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuCombineskuUpdateApiResults.Put(v)
}
