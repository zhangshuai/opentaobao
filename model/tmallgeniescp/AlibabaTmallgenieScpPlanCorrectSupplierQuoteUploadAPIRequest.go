package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest 供应商校准后的配额同步 API请求
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest 初始化AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) Reset() {
	r._netDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNetDemandRequest is NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
	r._netDemandRequest = _netDemandRequest
	r.Set("net_demand_request", _netDemandRequest)
	return nil
}

// GetNetDemandRequest NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
	return r._netDemandRequest
}

var poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest
func GetAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest 将 AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest(v *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest.Put(v)
}
