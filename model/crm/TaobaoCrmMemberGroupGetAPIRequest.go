package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberGroupGetAPIRequest 获取买家身上的标签 API请求
// taobao.crm.member.group.get
//
// 获取买家身上的标签，不返回标签的总人数
type TaobaoCrmMemberGroupGetAPIRequest struct {
	model.Params
	// 会员Nick
	_buyerNick string
}

// NewTaobaoCrmMemberGroupGetRequest 初始化TaobaoCrmMemberGroupGetAPIRequest对象
func NewTaobaoCrmMemberGroupGetRequest() *TaobaoCrmMemberGroupGetAPIRequest {
	return &TaobaoCrmMemberGroupGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMemberGroupGetAPIRequest) Reset() {
	r._buyerNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberGroupGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.member.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberGroupGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMemberGroupGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 会员Nick
func (r *TaobaoCrmMemberGroupGetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoCrmMemberGroupGetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

var poolTaobaoCrmMemberGroupGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMemberGroupGetRequest()
	},
}

// GetTaobaoCrmMemberGroupGetRequest 从 sync.Pool 获取 TaobaoCrmMemberGroupGetAPIRequest
func GetTaobaoCrmMemberGroupGetAPIRequest() *TaobaoCrmMemberGroupGetAPIRequest {
	return poolTaobaoCrmMemberGroupGetAPIRequest.Get().(*TaobaoCrmMemberGroupGetAPIRequest)
}

// ReleaseTaobaoCrmMemberGroupGetAPIRequest 将 TaobaoCrmMemberGroupGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMemberGroupGetAPIRequest(v *TaobaoCrmMemberGroupGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmMemberGroupGetAPIRequest.Put(v)
}
