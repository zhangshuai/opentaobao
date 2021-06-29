package alihealthmedical

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构通知平台"医生拒诊" API请求
alibaba.alihealth.medical.order.refuse

三方机构通知平台"医生拒诊"
*/
type AlibabaAlihealthMedicalOrderRefuseRequest struct {
    model.Params
    // 请求入参
    requestInfo   *RefuseOrderRequestDTO
}

// 初始化AlibabaAlihealthMedicalOrderRefuseRequest对象
func NewAlibabaAlihealthMedicalOrderRefuseRequest() *AlibabaAlihealthMedicalOrderRefuseRequest{
    return &AlibabaAlihealthMedicalOrderRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalOrderRefuseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medical.order.refuse"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalOrderRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestInfo Setter
// 请求入参
func (r *AlibabaAlihealthMedicalOrderRefuseRequest) SetRequestInfo(requestInfo *RefuseOrderRequestDTO) error {
    r.requestInfo = requestInfo
    r.Set("request_info", requestInfo)
    return nil
}

// RequestInfo Getter
func (r AlibabaAlihealthMedicalOrderRefuseRequest) GetRequestInfo() *RefuseOrderRequestDTO {
    return r.requestInfo
}
