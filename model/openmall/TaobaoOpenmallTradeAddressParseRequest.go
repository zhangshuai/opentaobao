package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall服务地址区域码解析 API请求
taobao.openmall.trade.address.parse

openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
*/
type TaobaoOpenmallTradeAddressParseRequest struct {
    model.Params
    // 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
    rawAddress   string
    // 渠道商分销者淘宝账号
    distributor   string
}

// 初始化TaobaoOpenmallTradeAddressParseRequest对象
func NewTaobaoOpenmallTradeAddressParseRequest() *TaobaoOpenmallTradeAddressParseRequest{
    return &TaobaoOpenmallTradeAddressParseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeAddressParseRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.address.parse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeAddressParseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RawAddress Setter
// 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
func (r *TaobaoOpenmallTradeAddressParseRequest) SetRawAddress(rawAddress string) error {
    r.rawAddress = rawAddress
    r.Set("raw_address", rawAddress)
    return nil
}

// RawAddress Getter
func (r TaobaoOpenmallTradeAddressParseRequest) GetRawAddress() string {
    return r.rawAddress
}
// Distributor Setter
// 渠道商分销者淘宝账号
func (r *TaobaoOpenmallTradeAddressParseRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeAddressParseRequest) GetDistributor() string {
    return r.distributor
}
