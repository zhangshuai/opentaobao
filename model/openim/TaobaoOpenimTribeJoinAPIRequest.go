package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeJoinAPIRequest OPENIM群主动加入 API请求
// taobao.openim.tribe.join
//
// OPENIM群主动加入
type TaobaoOpenimTribeJoinAPIRequest struct {
	model.Params
	// 用户信息
	_user *OpenImUser
	// 群id
	_tribeId int64
}

// NewTaobaoOpenimTribeJoinRequest 初始化TaobaoOpenimTribeJoinAPIRequest对象
func NewTaobaoOpenimTribeJoinRequest() *TaobaoOpenimTribeJoinAPIRequest {
	return &TaobaoOpenimTribeJoinAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeJoinAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.join"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeJoinAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeJoinAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeJoinAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeJoinAPIRequest) GetUser() *OpenImUser {
	return r._user
}

// SetTribeId is TribeId Setter
// 群id
func (r *TaobaoOpenimTribeJoinAPIRequest) SetTribeId(_tribeId int64) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoOpenimTribeJoinAPIRequest) GetTribeId() int64 {
	return r._tribeId
}
