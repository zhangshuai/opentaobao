package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增值服务订购服务验证 API请求
taobao.vas.service.validate

增值服务订购服务验证
*/
type TaobaoVasServiceValidateAPIRequest struct {
    model.Params
    // 服务编码
    _servCode   string
    // 用户昵称
    _nick   string
}

// 初始化TaobaoVasServiceValidateAPIRequest对象
func NewTaobaoVasServiceValidateRequest() *TaobaoVasServiceValidateAPIRequest{
    return &TaobaoVasServiceValidateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVasServiceValidateAPIRequest) GetApiMethodName() string {
    return "taobao.vas.service.validate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVasServiceValidateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServCode Setter
// 服务编码
func (r *TaobaoVasServiceValidateAPIRequest) SetServCode(_servCode string) error {
    r._servCode = _servCode
    r.Set("serv_code", _servCode)
    return nil
}

// ServCode Getter
func (r TaobaoVasServiceValidateAPIRequest) GetServCode() string {
    return r._servCode
}
// Nick Setter
// 用户昵称
func (r *TaobaoVasServiceValidateAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoVasServiceValidateAPIRequest) GetNick() string {
    return r._nick
}