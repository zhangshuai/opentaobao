package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinateprocessingAPIRequest 慧飞商家协同单处理完成接口 API请求
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
type AlitripagentcoordinateprocessingAPIRequest struct {
	model.Params
	// 协同单处理完成参数
	_processingDto *ProcessingDto
}

// NewAlitripagentcoordinateprocessingRequest 初始化AlitripagentcoordinateprocessingAPIRequest对象
func NewAlitripagentcoordinateprocessingRequest() *AlitripagentcoordinateprocessingAPIRequest {
	return &AlitripagentcoordinateprocessingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinateprocessingAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.processing"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinateprocessingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinateprocessingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProcessingDto is ProcessingDto Setter
// 协同单处理完成参数
func (r *AlitripagentcoordinateprocessingAPIRequest) SetProcessingDto(_processingDto *ProcessingDto) error {
	r._processingDto = _processingDto
	r.Set("processing_dto", _processingDto)
	return nil
}

// GetProcessingDto ProcessingDto Getter
func (r AlitripagentcoordinateprocessingAPIRequest) GetProcessingDto() *ProcessingDto {
	return r._processingDto
}