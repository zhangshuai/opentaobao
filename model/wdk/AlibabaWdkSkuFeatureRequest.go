package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品标记接口 API请求
alibaba.wdk.sku.feature

给淘鲜达商品属性之外的打标通用能力，满足商品一些特殊的需求，比如是否参加营销。
*/
type AlibabaWdkSkuFeatureRequest struct {
    model.Params
    // SkuFeatureDo
    param   *SkuFeatureDo
}

// 初始化AlibabaWdkSkuFeatureRequest对象
func NewAlibabaWdkSkuFeatureRequest() *AlibabaWdkSkuFeatureRequest{
    return &AlibabaWdkSkuFeatureRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuFeatureRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.feature"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuFeatureRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// SkuFeatureDo
func (r *AlibabaWdkSkuFeatureRequest) SetParam(param *SkuFeatureDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkSkuFeatureRequest) GetParam() *SkuFeatureDo {
    return r.param
}
