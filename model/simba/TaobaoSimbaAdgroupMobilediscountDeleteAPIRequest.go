package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest 批量删除adgroup的移动溢价 API请求
// taobao.simba.adgroup.mobilediscount.delete
//
// 批量删除adgroup的移动溢价
type TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest struct {
	model.Params
	// adgroup主键数组（批量最多支持200个）
	_adgroupIds []string
	// 昵称
	_nick string
}

// NewTaobaoSimbaAdgroupMobilediscountDeleteRequest 初始化TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest对象
func NewTaobaoSimbaAdgroupMobilediscountDeleteRequest() *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest {
	return &TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) Reset() {
	r._adgroupIds = r._adgroupIds[:0]
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.mobilediscount.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIds is AdgroupIds Setter
// adgroup主键数组（批量最多支持200个）
func (r *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) SetAdgroupIds(_adgroupIds []string) error {
	r._adgroupIds = _adgroupIds
	r.Set("adgroup_ids", _adgroupIds)
	return nil
}

// GetAdgroupIds AdgroupIds Getter
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetAdgroupIds() []string {
	return r._adgroupIds
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSimbaAdgroupMobilediscountDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupMobilediscountDeleteRequest()
	},
}

// GetTaobaoSimbaAdgroupMobilediscountDeleteRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest
func GetTaobaoSimbaAdgroupMobilediscountDeleteAPIRequest() *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest {
	return poolTaobaoSimbaAdgroupMobilediscountDeleteAPIRequest.Get().(*TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupMobilediscountDeleteAPIRequest 将 TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupMobilediscountDeleteAPIRequest(v *TaobaoSimbaAdgroupMobilediscountDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupMobilediscountDeleteAPIRequest.Put(v)
}
