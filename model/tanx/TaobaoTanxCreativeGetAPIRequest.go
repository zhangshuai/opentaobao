package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审核创意状态 API请求
taobao.tanx.creative.get

获取单个审核创意状态
*/
type TaobaoTanxCreativeGetAPIRequest struct {
    model.Params
    // DSP的memberId
    _memberId   int64
    // dsp用户身份认证的TOKEN
    _token   string
    // 当前时间戳，1970-01-01后的秒数
    _signTime   int64
    // 创意ID
    _creativeId   string
}

// 初始化TaobaoTanxCreativeGetAPIRequest对象
func NewTaobaoTanxCreativeGetRequest() *TaobaoTanxCreativeGetAPIRequest{
    return &TaobaoTanxCreativeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxCreativeGetAPIRequest) GetApiMethodName() string {
    return "taobao.tanx.creative.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxCreativeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxCreativeGetAPIRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxCreativeGetAPIRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativeGetAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxCreativeGetAPIRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativeGetAPIRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxCreativeGetAPIRequest) GetSignTime() int64 {
    return r._signTime
}
// CreativeId Setter
// 创意ID
func (r *TaobaoTanxCreativeGetAPIRequest) SetCreativeId(_creativeId string) error {
    r._creativeId = _creativeId
    r.Set("creative_id", _creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoTanxCreativeGetAPIRequest) GetCreativeId() string {
    return r._creativeId
}