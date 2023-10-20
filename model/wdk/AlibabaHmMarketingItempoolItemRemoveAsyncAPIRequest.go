package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolitemremoveasyncAPIRequest 商品池删除商品 API请求
// alibaba.hm.marketing.itempool.item.remove.async
//
// 新模型下删除商品
type AlibabahmmarketingitempoolitemremoveasyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemPoolSku
	// 活动信息
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabahmmarketingitempoolitemremoveasyncRequest 初始化AlibabahmmarketingitempoolitemremoveasyncAPIRequest对象
func NewAlibabahmmarketingitempoolitemremoveasyncRequest() *AlibabahmmarketingitempoolitemremoveasyncAPIRequest {
	return &AlibabahmmarketingitempoolitemremoveasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolitemremoveasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.item.remove.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolitemremoveasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolitemremoveasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabahmmarketingitempoolitemremoveasyncAPIRequest) SetParam0(_param0 []ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingitempoolitemremoveasyncAPIRequest) GetParam0() []ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动信息
func (r *AlibabahmmarketingitempoolitemremoveasyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingitempoolitemremoveasyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabahmmarketingitempoolitemremoveasyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabahmmarketingitempoolitemremoveasyncAPIRequest) GetVersion() int64 {
	return r._version
}