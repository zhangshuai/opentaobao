package icbulogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest 查询物流运力列表 API请求
// alibaba.onetouch.logistics.express.logistics.product.list
//
// 查询物流产品&amp;揽收仓库列表
type AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest struct {
	model.Params
}

// NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest 初始化AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest() *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.logistics.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressLogisticsProductListRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest
func GetAlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest() *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest 将 AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest(v *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest.Put(v)
}
