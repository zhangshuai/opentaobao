package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest 阿里健康新挂号数据回传 API请求
// alibaba.alihealth.medical.registration.syncnew
//
// 阿里健康新挂号记录回传接口
type AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest struct {
	model.Params
	// 接口入参
	_saveRequest *CommonRequest4top
}

// NewAlibabaAlihealthMedicalRegistrationSyncnewRequest 初始化AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncnewRequest() *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest {
	return &AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medical.registration.syncnew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSaveRequest is SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) SetSaveRequest(_saveRequest *CommonRequest4top) error {
	r._saveRequest = _saveRequest
	r.Set("save_request", _saveRequest)
	return nil
}

// GetSaveRequest SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest) GetSaveRequest() *CommonRequest4top {
	return r._saveRequest
}
