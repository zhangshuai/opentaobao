package fundplatform

import (
	"sync"
)

// AlibabaFundplatformCardorderMakeSuccessStruct 结构体
type AlibabaFundplatformCardorderMakeSuccessStruct struct {
	// 制卡时传入的外部订单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	OwnSign string `json:"own_sign,omitempty" xml:"own_sign,omitempty"`
	// 制卡单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolAlibabaFundplatformCardorderMakeSuccessStruct = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformCardorderMakeSuccessStruct)
	},
}

// GetAlibabaFundplatformCardorderMakeSuccessStruct() 从对象池中获取AlibabaFundplatformCardorderMakeSuccessStruct
func GetAlibabaFundplatformCardorderMakeSuccessStruct() *AlibabaFundplatformCardorderMakeSuccessStruct {
	return poolAlibabaFundplatformCardorderMakeSuccessStruct.Get().(*AlibabaFundplatformCardorderMakeSuccessStruct)
}

// ReleaseAlibabaFundplatformCardorderMakeSuccessStruct 释放AlibabaFundplatformCardorderMakeSuccessStruct
func ReleaseAlibabaFundplatformCardorderMakeSuccessStruct(v *AlibabaFundplatformCardorderMakeSuccessStruct) {
	v.OutBizId = ""
	v.OwnSign = ""
	v.OrderId = 0
	poolAlibabaFundplatformCardorderMakeSuccessStruct.Put(v)
}
