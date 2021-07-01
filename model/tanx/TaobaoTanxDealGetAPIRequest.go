package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外部dsp提供交易id查询接口 API请求
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
type TaobaoTanxDealGetAPIRequest struct {
    model.Params
    // dsp用户id
    _dspId   int64
    // 交易id
    _dealId   int64
    // 1970年到现在的时间，毫秒
    _signTime   int64
    // 验证token
    _token   string
}

// 初始化TaobaoTanxDealGetAPIRequest对象
func NewTaobaoTanxDealGetRequest() *TaobaoTanxDealGetAPIRequest{
    return &TaobaoTanxDealGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxDealGetAPIRequest) GetApiMethodName() string {
    return "taobao.tanx.deal.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxDealGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DspId Setter
// dsp用户id
func (r *TaobaoTanxDealGetAPIRequest) SetDspId(_dspId int64) error {
    r._dspId = _dspId
    r.Set("dsp_id", _dspId)
    return nil
}

// DspId Getter
func (r TaobaoTanxDealGetAPIRequest) GetDspId() int64 {
    return r._dspId
}
// DealId Setter
// 交易id
func (r *TaobaoTanxDealGetAPIRequest) SetDealId(_dealId int64) error {
    r._dealId = _dealId
    r.Set("deal_id", _dealId)
    return nil
}

// DealId Getter
func (r TaobaoTanxDealGetAPIRequest) GetDealId() int64 {
    return r._dealId
}
// SignTime Setter
// 1970年到现在的时间，毫秒
func (r *TaobaoTanxDealGetAPIRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxDealGetAPIRequest) GetSignTime() int64 {
    return r._signTime
}
// Token Setter
// 验证token
func (r *TaobaoTanxDealGetAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxDealGetAPIRequest) GetToken() string {
    return r._token
}