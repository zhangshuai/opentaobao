package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoskuscustomgetAPIRequest 根据外部ID取商品SKU API请求
// taobao.skus.custom.get
//
// 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku &lt;br/&gt;这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
type TaobaoskuscustomgetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
	_fields string
	// Sku的外部商家ID
	_outerId string
}

// NewTaobaoskuscustomgetRequest 初始化TaobaoskuscustomgetAPIRequest对象
func NewTaobaoskuscustomgetRequest() *TaobaoskuscustomgetAPIRequest {
	return &TaobaoskuscustomgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoskuscustomgetAPIRequest) GetApiMethodName() string {
	return "taobao.skus.custom.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoskuscustomgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoskuscustomgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
func (r *TaobaoskuscustomgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoskuscustomgetAPIRequest) GetFields() string {
	return r._fields
}

// SetOuterId is OuterId Setter
// Sku的外部商家ID
func (r *TaobaoskuscustomgetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoskuscustomgetAPIRequest) GetOuterId() string {
	return r._outerId
}
