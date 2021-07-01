package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询appstore应用订购关系 API请求
taobao.appstore.subscribe.get

查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
*/
type TaobaoAppstoreSubscribeGetAPIRequest struct {
    model.Params
    // 用户昵称
    _nick   string
}

// 初始化TaobaoAppstoreSubscribeGetAPIRequest对象
func NewTaobaoAppstoreSubscribeGetRequest() *TaobaoAppstoreSubscribeGetAPIRequest{
    return &TaobaoAppstoreSubscribeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppstoreSubscribeGetAPIRequest) GetApiMethodName() string {
    return "taobao.appstore.subscribe.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppstoreSubscribeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户昵称
func (r *TaobaoAppstoreSubscribeGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoAppstoreSubscribeGetAPIRequest) GetNick() string {
    return r._nick
}