package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品更新接口 API请求
alibaba.wdk.sku.combinesku.update

组合商品修改接口
*/
type AlibabaWdkSkuCombineskuUpdateRequest struct {
    model.Params
    // 请求参数
    paramList   []SkuDo
}

// 初始化AlibabaWdkSkuCombineskuUpdateRequest对象
func NewAlibabaWdkSkuCombineskuUpdateRequest() *AlibabaWdkSkuCombineskuUpdateRequest{
    return &AlibabaWdkSkuCombineskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuUpdateRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

// ParamList Getter
func (r AlibabaWdkSkuCombineskuUpdateRequest) GetParamList() []SkuDo {
    return r.paramList
}
