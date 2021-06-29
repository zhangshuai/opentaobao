package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商品特价活动 API请求
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动
*/
type AlibabaWdkMarketingItemdiscountCreateactivityRequest struct {
    model.Params
    // 创建活动请求入参
    param   *ItemDiscountActivityRequest
}

// 初始化AlibabaWdkMarketingItemdiscountCreateactivityRequest对象
func NewAlibabaWdkMarketingItemdiscountCreateactivityRequest() *AlibabaWdkMarketingItemdiscountCreateactivityRequest{
    return &AlibabaWdkMarketingItemdiscountCreateactivityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountCreateactivityRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itemdiscount.createactivity"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountCreateactivityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingItemdiscountCreateactivityRequest) SetParam(param *ItemDiscountActivityRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkMarketingItemdiscountCreateactivityRequest) GetParam() *ItemDiscountActivityRequest {
    return r.param
}
