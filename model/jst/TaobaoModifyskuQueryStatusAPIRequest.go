package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoModifyskuQueryStatusAPIRequest 查询商家是否开通自助修改商品信息服务 API请求
// taobao.modifysku.query.status
//
// 查询商家是否开通自助修改商品信息服务
// 1. 没有授权
// 2. 已与当前appkey签约
// 3. 没有签约
// 4. 与其他服务商软件签约，如果是同一个isv name，返回appkey，否则不返回。
type TaobaoModifyskuQueryStatusAPIRequest struct {
	model.Params
}

// NewTaobaoModifyskuQueryStatusRequest 初始化TaobaoModifyskuQueryStatusAPIRequest对象
func NewTaobaoModifyskuQueryStatusRequest() *TaobaoModifyskuQueryStatusAPIRequest {
	return &TaobaoModifyskuQueryStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoModifyskuQueryStatusAPIRequest) GetApiMethodName() string {
	return "taobao.modifysku.query.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoModifyskuQueryStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoModifyskuQueryStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}
