package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康新挂号数据回传 API请求
alibaba.alihealth.medical.registration.syncnew

阿里健康新挂号记录回传接口
*/
type AlibabaAlihealthMedicalRegistrationSyncnewRequest struct {
    model.Params
    // 接口入参
    saveRequest   *CommonRequest4Top
}

// 初始化AlibabaAlihealthMedicalRegistrationSyncnewRequest对象
func NewAlibabaAlihealthMedicalRegistrationSyncnewRequest() *AlibabaAlihealthMedicalRegistrationSyncnewRequest{
    return &AlibabaAlihealthMedicalRegistrationSyncnewRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalRegistrationSyncnewRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.registration.syncnew"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalRegistrationSyncnewRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveRequest Setter
// 接口入参
func (r *AlibabaAlihealthMedicalRegistrationSyncnewRequest) SetSaveRequest(saveRequest *CommonRequest4Top) error {
    r.saveRequest = saveRequest
    r.Set("save_request", saveRequest)
    return nil
}

// SaveRequest Getter
func (r AlibabaAlihealthMedicalRegistrationSyncnewRequest) GetSaveRequest() *CommonRequest4Top {
    return r.saveRequest
}