package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest 获取authcode API请求
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
type AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest struct {
	model.Params
}

// NewAlibabaAilabsTmallgenieAuthDeviceGetcodeRequest 初始化AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceGetcodeRequest() *AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.getcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}
