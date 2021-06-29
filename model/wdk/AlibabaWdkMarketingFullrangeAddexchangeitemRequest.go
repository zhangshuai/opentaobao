package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全场增加换购品 API请求
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品
*/
type AlibabaWdkMarketingFullrangeAddexchangeitemRequest struct {
    model.Params
    // 系统自动生成
    param0   *ItemStairSku
    // 系统自动生成
    param1   *CommonActivityParam
}

// 初始化AlibabaWdkMarketingFullrangeAddexchangeitemRequest对象
func NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest() *AlibabaWdkMarketingFullrangeAddexchangeitemRequest{
    return &AlibabaWdkMarketingFullrangeAddexchangeitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetApiMethodName() string {
    return "alibaba.wdk.marketing.fullrange.addexchangeitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemRequest) SetParam0(param0 *ItemStairSku) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetParam0() *ItemStairSku {
    return r.param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemRequest) SetParam1(param1 *CommonActivityParam) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemRequest) GetParam1() *CommonActivityParam {
    return r.param1
}
