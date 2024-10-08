package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionUpdateAPIRequest 更新商品分销内容 API请求
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
type AlibabaDchainAoxiangItemDistributionUpdateAPIRequest struct {
	model.Params
	// 更新商品分销入参
	_updateItemDistributionRequest *UpdateItemDistributionRequest
}

// NewAlibabaDchainAoxiangItemDistributionUpdateRequest 初始化AlibabaDchainAoxiangItemDistributionUpdateAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionUpdateRequest() *AlibabaDchainAoxiangItemDistributionUpdateAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) Reset() {
	r._updateItemDistributionRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateItemDistributionRequest is UpdateItemDistributionRequest Setter
// 更新商品分销入参
func (r *AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) SetUpdateItemDistributionRequest(_updateItemDistributionRequest *UpdateItemDistributionRequest) error {
	r._updateItemDistributionRequest = _updateItemDistributionRequest
	r.Set("update_item_distribution_request", _updateItemDistributionRequest)
	return nil
}

// GetUpdateItemDistributionRequest UpdateItemDistributionRequest Getter
func (r AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) GetUpdateItemDistributionRequest() *UpdateItemDistributionRequest {
	return r._updateItemDistributionRequest
}

var poolAlibabaDchainAoxiangItemDistributionUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemDistributionUpdateRequest()
	},
}

// GetAlibabaDchainAoxiangItemDistributionUpdateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionUpdateAPIRequest
func GetAlibabaDchainAoxiangItemDistributionUpdateAPIRequest() *AlibabaDchainAoxiangItemDistributionUpdateAPIRequest {
	return poolAlibabaDchainAoxiangItemDistributionUpdateAPIRequest.Get().(*AlibabaDchainAoxiangItemDistributionUpdateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemDistributionUpdateAPIRequest 将 AlibabaDchainAoxiangItemDistributionUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionUpdateAPIRequest(v *AlibabaDchainAoxiangItemDistributionUpdateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionUpdateAPIRequest.Put(v)
}
