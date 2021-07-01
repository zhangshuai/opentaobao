package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取认证申请信息 API请求
alibaba.nazca.token.authapply.get

根据token获取认证申请信息
*/
type AlibabaNazcaTokenAuthapplyGetAPIRequest struct {
    model.Params
    // token
    _token   string
}

// 初始化AlibabaNazcaTokenAuthapplyGetAPIRequest对象
func NewAlibabaNazcaTokenAuthapplyGetRequest() *AlibabaNazcaTokenAuthapplyGetAPIRequest{
    return &AlibabaNazcaTokenAuthapplyGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.authapply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// token
func (r *AlibabaNazcaTokenAuthapplyGetAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaNazcaTokenAuthapplyGetAPIRequest) GetToken() string {
    return r._token
}