package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappriskdetailgetAPIRequest 应用风险详细信息查询接口 API请求
// alibaba.security.jaq.app.riskdetail.get
//
// 用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
type AlibabasecurityjaqappriskdetailgetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
	// 本地化语言信息
	_locale *Locale
}

// NewAlibabasecurityjaqappriskdetailgetRequest 初始化AlibabasecurityjaqappriskdetailgetAPIRequest对象
func NewAlibabasecurityjaqappriskdetailgetRequest() *AlibabasecurityjaqappriskdetailgetAPIRequest {
	return &AlibabasecurityjaqappriskdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqappriskdetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.riskdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqappriskdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqappriskdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 任务唯一标识
func (r *AlibabasecurityjaqappriskdetailgetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabasecurityjaqappriskdetailgetAPIRequest) GetItemId() string {
	return r._itemId
}

// SetLocale is Locale Setter
// 本地化语言信息
func (r *AlibabasecurityjaqappriskdetailgetAPIRequest) SetLocale(_locale *Locale) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AlibabasecurityjaqappriskdetailgetAPIRequest) GetLocale() *Locale {
	return r._locale
}
