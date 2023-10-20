package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest 交易猫外部商家商品状态批量查询接口 API请求
// alibaba.jym.item.external.goods.status.batch.query
//
// 供外部B端商家接入，请求查询商品状态，返回商品状态查询结果
type AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest struct {
	model.Params
	// 批量查询商品状态接口入参
	_batchGoodsStatusQuery *BatchGoodsStatusQueryDto
}

// NewAlibabaJymItemExternalGoodsStatusBatchQueryRequest 初始化AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest对象
func NewAlibabaJymItemExternalGoodsStatusBatchQueryRequest() *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest {
	return &AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) Reset() {
	r._batchGoodsStatusQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.status.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchGoodsStatusQuery is BatchGoodsStatusQuery Setter
// 批量查询商品状态接口入参
func (r *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) SetBatchGoodsStatusQuery(_batchGoodsStatusQuery *BatchGoodsStatusQueryDto) error {
	r._batchGoodsStatusQuery = _batchGoodsStatusQuery
	r.Set("batch_goods_status_query", _batchGoodsStatusQuery)
	return nil
}

// GetBatchGoodsStatusQuery BatchGoodsStatusQuery Getter
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetBatchGoodsStatusQuery() *BatchGoodsStatusQueryDto {
	return r._batchGoodsStatusQuery
}

var poolAlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymItemExternalGoodsStatusBatchQueryRequest()
	},
}

// GetAlibabaJymItemExternalGoodsStatusBatchQueryRequest 从 sync.Pool 获取 AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest
func GetAlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest() *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest {
	return poolAlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest.Get().(*AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest)
}

// ReleaseAlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest 将 AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest(v *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest.Put(v)
}
