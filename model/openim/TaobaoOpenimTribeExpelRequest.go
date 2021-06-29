package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群踢出成员 API请求
taobao.openim.tribe.expel

OPENIM群踢出成员
*/
type TaobaoOpenimTribeExpelRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群id
    tribeId   int64
    // 用户信息
    member   *OpenImUser
}

// 初始化TaobaoOpenimTribeExpelRequest对象
func NewTaobaoOpenimTribeExpelRequest() *TaobaoOpenimTribeExpelRequest{
    return &TaobaoOpenimTribeExpelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeExpelRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.expel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeExpelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeExpelRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeExpelRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeExpelRequest) GetTribeId() int64 {
    return r.tribeId
}
// Member Setter
// 用户信息
func (r *TaobaoOpenimTribeExpelRequest) SetMember(member *OpenImUser) error {
    r.member = member
    r.Set("member", member)
    return nil
}

// Member Getter
func (r TaobaoOpenimTribeExpelRequest) GetMember() *OpenImUser {
    return r.member
}
