package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegistrationSyncAPIRequest 阿里健康支付宝挂号记录回传接口 API请求
// alibaba.alihealth.medical.registration.sync
//
// 阿里健康支付宝挂号记录回传接口
type AlibabaAlihealthMedicalRegistrationSyncAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4top
}

// NewAlibabaAlihealthMedicalRegistrationSyncRequest 初始化AlibabaAlihealthMedicalRegistrationSyncAPIRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncRequest() *AlibabaAlihealthMedicalRegistrationSyncAPIRequest {
	return &AlibabaAlihealthMedicalRegistrationSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.registration.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncAPIRequest) GetSaveRequest() *CommonRequest4top {
	return r._saveRequest
}
