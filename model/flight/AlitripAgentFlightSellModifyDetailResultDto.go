package flight

import (
	"sync"
)

// AlitripAgentFlightSellModifyDetailResultDto 结构体
type AlitripAgentFlightSellModifyDetailResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果对象
	Data *ModifyDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellModifyDetailResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyDetailResultDto)
	},
}

// GetAlitripAgentFlightSellModifyDetailResultDto() 从对象池中获取AlitripAgentFlightSellModifyDetailResultDto
func GetAlitripAgentFlightSellModifyDetailResultDto() *AlitripAgentFlightSellModifyDetailResultDto {
	return poolAlitripAgentFlightSellModifyDetailResultDto.Get().(*AlitripAgentFlightSellModifyDetailResultDto)
}

// ReleaseAlitripAgentFlightSellModifyDetailResultDto 释放AlitripAgentFlightSellModifyDetailResultDto
func ReleaseAlitripAgentFlightSellModifyDetailResultDto(v *AlitripAgentFlightSellModifyDetailResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripAgentFlightSellModifyDetailResultDto.Put(v)
}
