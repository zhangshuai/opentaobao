package tmallgeniescp

import (
	"sync"
)

// DataResult 结构体
type DataResult struct {
	// 数据对象列表
	DataList []IbpChannelDto `json:"data_list,omitempty" xml:"data_list>ibp_channel_dto,omitempty"`
	// 结果msg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 结果code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 参数msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 参数code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolDataResult = sync.Pool{
	New: func() any {
		return new(DataResult)
	},
}

// GetDataResult() 从对象池中获取DataResult
func GetDataResult() *DataResult {
	return poolDataResult.Get().(*DataResult)
}

// ReleaseDataResult 释放DataResult
func ReleaseDataResult(v *DataResult) {
	v.DataList = v.DataList[:0]
	v.ResultMsg = ""
	v.TraceId = ""
	v.ResultCode = ""
	v.Msg = ""
	v.Code = ""
	poolDataResult.Put(v)
}
