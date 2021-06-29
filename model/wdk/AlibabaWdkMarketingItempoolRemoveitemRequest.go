package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
移除商品池里面的商品 API请求
alibaba.wdk.marketing.itempool.removeitem

移除商品池里面的商品
*/
type AlibabaWdkMarketingItempoolRemoveitemRequest struct {
    model.Params
    // 商品对象
    param0   *ItemPoolSku
    // 活动基本信息
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingItempoolRemoveitemRequest对象
func NewAlibabaWdkMarketingItempoolRemoveitemRequest() *AlibabaWdkMarketingItempoolRemoveitemRequest{
    return &AlibabaWdkMarketingItempoolRemoveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.itempool.removeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolRemoveitemRequest) SetParam0(param0 *ItemPoolSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetParam0() *ItemPoolSku {
    return r.param0
}
// Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolRemoveitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingItempoolRemoveitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
