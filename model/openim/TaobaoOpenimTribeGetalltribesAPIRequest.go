package openim

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGetalltribesAPIRequest 获取用户群列表 API请求
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
type TaobaoOpenimTribeGetalltribesAPIRequest struct {
	model.Params
	// 群类型
	_tribeTypes []int64
	// 用户信息
	_user *OpenImUser
}

// NewTaobaoOpenimTribeGetalltribesRequest 初始化TaobaoOpenimTribeGetalltribesAPIRequest对象
func NewTaobaoOpenimTribeGetalltribesRequest() *TaobaoOpenimTribeGetalltribesAPIRequest {
	return &TaobaoOpenimTribeGetalltribesAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenimTribeGetalltribesAPIRequest) Reset() {
	r._tribeTypes = r._tribeTypes[:0]
	r._user = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribe.getalltribes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTribeTypes is TribeTypes Setter
// 群类型
func (r *TaobaoOpenimTribeGetalltribesAPIRequest) SetTribeTypes(_tribeTypes []int64) error {
	r._tribeTypes = _tribeTypes
	r.Set("tribe_types", _tribeTypes)
	return nil
}

// GetTribeTypes TribeTypes Getter
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetTribeTypes() []int64 {
	return r._tribeTypes
}

// SetUser is User Setter
// 用户信息
func (r *TaobaoOpenimTribeGetalltribesAPIRequest) SetUser(_user *OpenImUser) error {
	r._user = _user
	r.Set("user", _user)
	return nil
}

// GetUser User Getter
func (r TaobaoOpenimTribeGetalltribesAPIRequest) GetUser() *OpenImUser {
	return r._user
}

var poolTaobaoOpenimTribeGetalltribesAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenimTribeGetalltribesRequest()
	},
}

// GetTaobaoOpenimTribeGetalltribesRequest 从 sync.Pool 获取 TaobaoOpenimTribeGetalltribesAPIRequest
func GetTaobaoOpenimTribeGetalltribesAPIRequest() *TaobaoOpenimTribeGetalltribesAPIRequest {
	return poolTaobaoOpenimTribeGetalltribesAPIRequest.Get().(*TaobaoOpenimTribeGetalltribesAPIRequest)
}

// ReleaseTaobaoOpenimTribeGetalltribesAPIRequest 将 TaobaoOpenimTribeGetalltribesAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenimTribeGetalltribesAPIRequest(v *TaobaoOpenimTribeGetalltribesAPIRequest) {
	v.Reset()
	poolTaobaoOpenimTribeGetalltribesAPIRequest.Put(v)
}
