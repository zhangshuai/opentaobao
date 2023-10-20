package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemtemplatesgetAPIRequest 获取用户宝贝详情页模板名称 API请求
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoitemtemplatesgetAPIRequest struct {
	model.Params
}

// NewTaobaoitemtemplatesgetRequest 初始化TaobaoitemtemplatesgetAPIRequest对象
func NewTaobaoitemtemplatesgetRequest() *TaobaoitemtemplatesgetAPIRequest {
	return &TaobaoitemtemplatesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemtemplatesgetAPIRequest) GetApiMethodName() string {
	return "taobao.item.templates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemtemplatesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemtemplatesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}