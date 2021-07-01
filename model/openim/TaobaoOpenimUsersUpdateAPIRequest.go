package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新用户信息 API请求
taobao.openim.users.update

批量更新用户信息
*/
type TaobaoOpenimUsersUpdateAPIRequest struct {
    model.Params
    // 用户信息列表
    _userinfos   []Userinfos
}

// 初始化TaobaoOpenimUsersUpdateAPIRequest对象
func NewTaobaoOpenimUsersUpdateRequest() *TaobaoOpenimUsersUpdateAPIRequest{
    return &TaobaoOpenimUsersUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimUsersUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.openim.users.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimUsersUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Userinfos Setter
// 用户信息列表
func (r *TaobaoOpenimUsersUpdateAPIRequest) SetUserinfos(_userinfos []Userinfos) error {
    r._userinfos = _userinfos
    r.Set("userinfos", _userinfos)
    return nil
}

// Userinfos Getter
func (r TaobaoOpenimUsersUpdateAPIRequest) GetUserinfos() []Userinfos {
    return r._userinfos
}