package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarxcarsynchronizecarmodeldataAPIRequest 爱车车型数据同步 API请求
// tmall.car.xcar.synchronize.car.model.data
//
// 爱车汽车车型数据同步到天猫
type TmallcarxcarsynchronizecarmodeldataAPIRequest struct {
	model.Params
	// 传入对象描述
	_paramXCarSysModelDTO *XcarSysModelDto
}

// NewTmallcarxcarsynchronizecarmodeldataRequest 初始化TmallcarxcarsynchronizecarmodeldataAPIRequest对象
func NewTmallcarxcarsynchronizecarmodeldataRequest() *TmallcarxcarsynchronizecarmodeldataAPIRequest {
	return &TmallcarxcarsynchronizecarmodeldataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarxcarsynchronizecarmodeldataAPIRequest) GetApiMethodName() string {
	return "tmall.car.xcar.synchronize.car.model.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarxcarsynchronizecarmodeldataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarxcarsynchronizecarmodeldataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamXCarSysModelDTO is ParamXCarSysModelDTO Setter
// 传入对象描述
func (r *TmallcarxcarsynchronizecarmodeldataAPIRequest) SetParamXCarSysModelDTO(_paramXCarSysModelDTO *XcarSysModelDto) error {
	r._paramXCarSysModelDTO = _paramXCarSysModelDTO
	r.Set("param_x_car_sys_model_d_t_o", _paramXCarSysModelDTO)
	return nil
}

// GetParamXCarSysModelDTO ParamXCarSysModelDTO Getter
func (r TmallcarxcarsynchronizecarmodeldataAPIRequest) GetParamXCarSysModelDTO() *XcarSysModelDto {
	return r._paramXCarSysModelDTO
}
