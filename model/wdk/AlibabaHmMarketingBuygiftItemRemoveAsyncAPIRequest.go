package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingbuygiftitemremoveasyncAPIRequest 批量删除买赠商品 API请求
// alibaba.hm.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
type AlibabahmmarketingbuygiftitemremoveasyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemBuyGiftSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabahmmarketingbuygiftitemremoveasyncRequest 初始化AlibabahmmarketingbuygiftitemremoveasyncAPIRequest对象
func NewAlibabahmmarketingbuygiftitemremoveasyncRequest() *AlibabahmmarketingbuygiftitemremoveasyncAPIRequest {
	return &AlibabahmmarketingbuygiftitemremoveasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.buygift.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) SetParam0(_param0 []ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) GetParam0() []ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabahmmarketingbuygiftitemremoveasyncAPIRequest) GetVersion() int64 {
	return r._version
}