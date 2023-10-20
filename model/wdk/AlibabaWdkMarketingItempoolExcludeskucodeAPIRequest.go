package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolexcludeskucodeAPIRequest 商品池排除商品【品类优惠使用】 API请求
// alibaba.wdk.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
type AlibabawdkmarketingitempoolexcludeskucodeAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingitempoolexcludeskucodeRequest 初始化AlibabawdkmarketingitempoolexcludeskucodeAPIRequest对象
func NewAlibabawdkmarketingitempoolexcludeskucodeRequest() *AlibabawdkmarketingitempoolexcludeskucodeAPIRequest {
	return &AlibabawdkmarketingitempoolexcludeskucodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.excludeskucode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingitempoolexcludeskucodeAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
