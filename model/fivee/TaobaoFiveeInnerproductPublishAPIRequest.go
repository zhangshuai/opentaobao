package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国产商品发布 API请求
taobao.fivee.innerproduct.publish

资质共享平台国产商品发布
*/
type TaobaoFiveeInnerproductPublishAPIRequest struct {
    model.Params
    // bu身份标识
    _paramBucode   string
    // 国产商品
    _paramInnerProduct   *InnerProduct
}

// 初始化TaobaoFiveeInnerproductPublishAPIRequest对象
func NewTaobaoFiveeInnerproductPublishRequest() *TaobaoFiveeInnerproductPublishAPIRequest{
    return &TaobaoFiveeInnerproductPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetApiMethodName() string {
    return "taobao.fivee.innerproduct.publish"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeInnerproductPublishAPIRequest) SetParamBucode(_paramBucode string) error {
    r._paramBucode = _paramBucode
    r.Set("param_bucode", _paramBucode)
    return nil
}

// ParamBucode Getter
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetParamBucode() string {
    return r._paramBucode
}
// ParamInnerProduct Setter
// 国产商品
func (r *TaobaoFiveeInnerproductPublishAPIRequest) SetParamInnerProduct(_paramInnerProduct *InnerProduct) error {
    r._paramInnerProduct = _paramInnerProduct
    r.Set("param_inner_product", _paramInnerProduct)
    return nil
}

// ParamInnerProduct Getter
func (r TaobaoFiveeInnerproductPublishAPIRequest) GetParamInnerProduct() *InnerProduct {
    return r._paramInnerProduct
}