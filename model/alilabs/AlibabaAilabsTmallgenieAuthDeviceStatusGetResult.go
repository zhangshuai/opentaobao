package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceStatusGetResult struct {
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// IP地址
	DeviceIp string `json:"device_ip,omitempty" xml:"device_ip,omitempty"`
	// 扩展返回，保留使用
	Extensions string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// 在线状态（0：离线，1：在线）
	OnlineStatus string `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 设备ID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsTmallgenieAuthDeviceStatusGetResult `json:"result,omitempty" xml:"result,omitempty"`
	// 状态码（200：成功，其他：失败）
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthDeviceStatusGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceStatusGetResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceStatusGetResult() 从对象池中获取AlibabaAilabsTmallgenieAuthDeviceStatusGetResult
func GetAlibabaAilabsTmallgenieAuthDeviceStatusGetResult() *AlibabaAilabsTmallgenieAuthDeviceStatusGetResult {
	return poolAlibabaAilabsTmallgenieAuthDeviceStatusGetResult.Get().(*AlibabaAilabsTmallgenieAuthDeviceStatusGetResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetResult 释放AlibabaAilabsTmallgenieAuthDeviceStatusGetResult
func ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetResult(v *AlibabaAilabsTmallgenieAuthDeviceStatusGetResult) {
	v.Message = ""
	v.DeviceIp = ""
	v.Extensions = ""
	v.OnlineStatus = ""
	v.Uuid = ""
	v.Result = nil
	v.Code = 0
	poolAlibabaAilabsTmallgenieAuthDeviceStatusGetResult.Put(v)
}
