package flight

import (
	"sync"
)

// AlitripAgentFlightSellRefundApproveResultDto 结构体
type AlitripAgentFlightSellRefundApproveResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellRefundApproveResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellRefundApproveResultDto)
	},
}

// GetAlitripAgentFlightSellRefundApproveResultDto() 从对象池中获取AlitripAgentFlightSellRefundApproveResultDto
func GetAlitripAgentFlightSellRefundApproveResultDto() *AlitripAgentFlightSellRefundApproveResultDto {
	return poolAlitripAgentFlightSellRefundApproveResultDto.Get().(*AlitripAgentFlightSellRefundApproveResultDto)
}

// ReleaseAlitripAgentFlightSellRefundApproveResultDto 释放AlitripAgentFlightSellRefundApproveResultDto
func ReleaseAlitripAgentFlightSellRefundApproveResultDto(v *AlitripAgentFlightSellRefundApproveResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentFlightSellRefundApproveResultDto.Put(v)
}
