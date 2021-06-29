package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alios cosmo服务调用 API请求
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口
*/
type AliyunCosmoGatewayInvokeRequest struct {
    model.Params
    // 请求上下文参数
    context   *RdamContext
    // 请求对象
    rdamRequest   *RdamGenericRequest
}

// 初始化AliyunCosmoGatewayInvokeRequest对象
func NewAliyunCosmoGatewayInvokeRequest() *AliyunCosmoGatewayInvokeRequest{
    return &AliyunCosmoGatewayInvokeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunCosmoGatewayInvokeRequest) GetApiMethodName() string {
    return "aliyun.cosmo.gateway.invoke"
}

// IRequest interface 方法, 获取API参数
func (r AliyunCosmoGatewayInvokeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 请求上下文参数
func (r *AliyunCosmoGatewayInvokeRequest) SetContext(context *RdamContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AliyunCosmoGatewayInvokeRequest) GetContext() *RdamContext {
    return r.context
}
// RdamRequest Setter
// 请求对象
func (r *AliyunCosmoGatewayInvokeRequest) SetRdamRequest(rdamRequest *RdamGenericRequest) error {
    r.rdamRequest = rdamRequest
    r.Set("rdam_request", rdamRequest)
    return nil
}

// RdamRequest Getter
func (r AliyunCosmoGatewayInvokeRequest) GetRdamRequest() *RdamGenericRequest {
    return r.rdamRequest
}
