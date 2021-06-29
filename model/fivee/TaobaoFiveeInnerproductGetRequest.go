package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国产商品资质查询 API请求
taobao.fivee.innerproduct.get

资质共享平台，国产商品查询
*/
type TaobaoFiveeInnerproductGetRequest struct {
    model.Params
    // bu身份标识
    paramBucode   string
    // 条形码
    paramBarcode   string
}

// 初始化TaobaoFiveeInnerproductGetRequest对象
func NewTaobaoFiveeInnerproductGetRequest() *TaobaoFiveeInnerproductGetRequest{
    return &TaobaoFiveeInnerproductGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFiveeInnerproductGetRequest) GetApiMethodName() string {
    return "taobao.fivee.innerproduct.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFiveeInnerproductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBucode Setter
// bu身份标识
func (r *TaobaoFiveeInnerproductGetRequest) SetParamBucode(paramBucode string) error {
    r.paramBucode = paramBucode
    r.Set("param_bucode", paramBucode)
    return nil
}

// ParamBucode Getter
func (r TaobaoFiveeInnerproductGetRequest) GetParamBucode() string {
    return r.paramBucode
}
// ParamBarcode Setter
// 条形码
func (r *TaobaoFiveeInnerproductGetRequest) SetParamBarcode(paramBarcode string) error {
    r.paramBarcode = paramBarcode
    r.Set("param_barcode", paramBarcode)
    return nil
}

// ParamBarcode Getter
func (r TaobaoFiveeInnerproductGetRequest) GetParamBarcode() string {
    return r.paramBarcode
}