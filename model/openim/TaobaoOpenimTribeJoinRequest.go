package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OPENIM群主动加入 API请求
taobao.openim.tribe.join

OPENIM群主动加入
*/
type TaobaoOpenimTribeJoinRequest struct {
    model.Params
    // 用户信息
    user   *OpenImUser
    // 群id
    tribeId   int64
}

// 初始化TaobaoOpenimTribeJoinRequest对象
func NewTaobaoOpenimTribeJoinRequest() *TaobaoOpenimTribeJoinRequest{
    return &TaobaoOpenimTribeJoinRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeJoinRequest) GetApiMethodName() string {
    return "taobao.openim.tribe.join"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeJoinRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// User Setter
// 用户信息
func (r *TaobaoOpenimTribeJoinRequest) SetUser(user *OpenImUser) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r TaobaoOpenimTribeJoinRequest) GetUser() *OpenImUser {
    return r.user
}
// TribeId Setter
// 群id
func (r *TaobaoOpenimTribeJoinRequest) SetTribeId(tribeId int64) error {
    r.tribeId = tribeId
    r.Set("tribe_id", tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribeJoinRequest) GetTribeId() int64 {
    return r.tribeId
}
