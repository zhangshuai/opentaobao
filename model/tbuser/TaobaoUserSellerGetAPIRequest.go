package tbuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousersellergetAPIRequest 查询卖家用户信息 API请求
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
type TaobaousersellergetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
	_fields []string
}

// NewTaobaousersellergetRequest 初始化TaobaousersellergetAPIRequest对象
func NewTaobaousersellergetRequest() *TaobaousersellergetAPIRequest {
	return &TaobaousersellergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousersellergetAPIRequest) GetApiMethodName() string {
	return "taobao.user.seller.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousersellergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousersellergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
func (r *TaobaousersellergetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaousersellergetAPIRequest) GetFields() []string {
	return r._fields
}