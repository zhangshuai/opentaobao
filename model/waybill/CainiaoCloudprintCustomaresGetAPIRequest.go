package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintcustomaresgetAPIRequest 获取商家的自定义区模板信息 API请求
// cainiao.cloudprint.customares.get
//
// 供isv使用，获取商家的自定义区的模板信息
type CainiaocloudprintcustomaresgetAPIRequest struct {
	model.Params
	// 用户使用的标准模板id
	_templateId int64
}

// NewCainiaocloudprintcustomaresgetRequest 初始化CainiaocloudprintcustomaresgetAPIRequest对象
func NewCainiaocloudprintcustomaresgetRequest() *CainiaocloudprintcustomaresgetAPIRequest {
	return &CainiaocloudprintcustomaresgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocloudprintcustomaresgetAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.customares.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocloudprintcustomaresgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocloudprintcustomaresgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTemplateId is TemplateId Setter
// 用户使用的标准模板id
func (r *CainiaocloudprintcustomaresgetAPIRequest) SetTemplateId(_templateId int64) error {
	r._templateId = _templateId
	r.Set("template_id", _templateId)
	return nil
}

// GetTemplateId TemplateId Getter
func (r CainiaocloudprintcustomaresgetAPIRequest) GetTemplateId() int64 {
	return r._templateId
}
