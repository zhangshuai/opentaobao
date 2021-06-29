package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）ICBU商品发布schema接口 API请求
alibaba.icbu.product.schema.get

获取ICBU商品发布的页面规则和填写字段，适用于新发商品
*/
type AlibabaIcbuProductSchemaGetRequest struct {
    model.Params
    // 商品规则渲染请求
    paramProductTopPublishRequest   *ProductTopPublishRequest
}

// 初始化AlibabaIcbuProductSchemaGetRequest对象
func NewAlibabaIcbuProductSchemaGetRequest() *AlibabaIcbuProductSchemaGetRequest{
    return &AlibabaIcbuProductSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamProductTopPublishRequest Setter
// 商品规则渲染请求
func (r *AlibabaIcbuProductSchemaGetRequest) SetParamProductTopPublishRequest(paramProductTopPublishRequest *ProductTopPublishRequest) error {
    r.paramProductTopPublishRequest = paramProductTopPublishRequest
    r.Set("param_product_top_publish_request", paramProductTopPublishRequest)
    return nil
}

// ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaGetRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
    return r.paramProductTopPublishRequest
}
