package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销涉及到的商品的库存同步接口 API请求
alibaba.wdk.bm.stock.publish

用于操作sku的库存
*/
type AlibabaWdkBmStockPublishRequest struct {
    model.Params
    // 批量入参
    skuStockPublishParamList   []SkuStockPublishParamDo
}

// 初始化AlibabaWdkBmStockPublishRequest对象
func NewAlibabaWdkBmStockPublishRequest() *AlibabaWdkBmStockPublishRequest{
    return &AlibabaWdkBmStockPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmStockPublishRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.stock.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmStockPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkuStockPublishParamList Setter
// 批量入参
func (r *AlibabaWdkBmStockPublishRequest) SetSkuStockPublishParamList(skuStockPublishParamList []SkuStockPublishParamDo) error {
    r.skuStockPublishParamList = skuStockPublishParamList
    r.Set("sku_stock_publish_param_list", skuStockPublishParamList)
    return nil
}

// SkuStockPublishParamList Getter
func (r AlibabaWdkBmStockPublishRequest) GetSkuStockPublishParamList() []SkuStockPublishParamDo {
    return r.skuStockPublishParamList
}
