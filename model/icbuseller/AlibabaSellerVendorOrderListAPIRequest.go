package icbuseller

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorOrderListAPIRequest 国际站服务市场订单列表接口 API请求
// alibaba.seller.vendor.order.list
//
// 返回服务商在服务市场的客户订单
type AlibabaSellerVendorOrderListAPIRequest struct {
	model.Params
	// 查询参数
	_queryTradeDto *QueryTradeDto
}

// NewAlibabaSellerVendorOrderListRequest 初始化AlibabaSellerVendorOrderListAPIRequest对象
func NewAlibabaSellerVendorOrderListRequest() *AlibabaSellerVendorOrderListAPIRequest {
	return &AlibabaSellerVendorOrderListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSellerVendorOrderListAPIRequest) Reset() {
	r._queryTradeDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorOrderListAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.order.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorOrderListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSellerVendorOrderListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryTradeDto is QueryTradeDto Setter
// 查询参数
func (r *AlibabaSellerVendorOrderListAPIRequest) SetQueryTradeDto(_queryTradeDto *QueryTradeDto) error {
	r._queryTradeDto = _queryTradeDto
	r.Set("query_trade_dto", _queryTradeDto)
	return nil
}

// GetQueryTradeDto QueryTradeDto Getter
func (r AlibabaSellerVendorOrderListAPIRequest) GetQueryTradeDto() *QueryTradeDto {
	return r._queryTradeDto
}

var poolAlibabaSellerVendorOrderListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSellerVendorOrderListRequest()
	},
}

// GetAlibabaSellerVendorOrderListRequest 从 sync.Pool 获取 AlibabaSellerVendorOrderListAPIRequest
func GetAlibabaSellerVendorOrderListAPIRequest() *AlibabaSellerVendorOrderListAPIRequest {
	return poolAlibabaSellerVendorOrderListAPIRequest.Get().(*AlibabaSellerVendorOrderListAPIRequest)
}

// ReleaseAlibabaSellerVendorOrderListAPIRequest 将 AlibabaSellerVendorOrderListAPIRequest 放入 sync.Pool
func ReleaseAlibabaSellerVendorOrderListAPIRequest(v *AlibabaSellerVendorOrderListAPIRequest) {
	v.Reset()
	poolAlibabaSellerVendorOrderListAPIRequest.Put(v)
}
