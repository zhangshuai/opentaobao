package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsCpGetAPIRequest 快递公司资源列表查询接口 API请求
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
type AlibabaAscpLogisticsCpGetAPIRequest struct {
	model.Params
	// 请求体
	_logisticsResourceRequest *LogisticsResourceRequest
}

// NewAlibabaAscpLogisticsCpGetRequest 初始化AlibabaAscpLogisticsCpGetAPIRequest对象
func NewAlibabaAscpLogisticsCpGetRequest() *AlibabaAscpLogisticsCpGetAPIRequest {
	return &AlibabaAscpLogisticsCpGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsCpGetAPIRequest) Reset() {
	r._logisticsResourceRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsCpGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.cp.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsCpGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsCpGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsResourceRequest is LogisticsResourceRequest Setter
// 请求体
func (r *AlibabaAscpLogisticsCpGetAPIRequest) SetLogisticsResourceRequest(_logisticsResourceRequest *LogisticsResourceRequest) error {
	r._logisticsResourceRequest = _logisticsResourceRequest
	r.Set("logistics_resource_request", _logisticsResourceRequest)
	return nil
}

// GetLogisticsResourceRequest LogisticsResourceRequest Getter
func (r AlibabaAscpLogisticsCpGetAPIRequest) GetLogisticsResourceRequest() *LogisticsResourceRequest {
	return r._logisticsResourceRequest
}

var poolAlibabaAscpLogisticsCpGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsCpGetRequest()
	},
}

// GetAlibabaAscpLogisticsCpGetRequest 从 sync.Pool 获取 AlibabaAscpLogisticsCpGetAPIRequest
func GetAlibabaAscpLogisticsCpGetAPIRequest() *AlibabaAscpLogisticsCpGetAPIRequest {
	return poolAlibabaAscpLogisticsCpGetAPIRequest.Get().(*AlibabaAscpLogisticsCpGetAPIRequest)
}

// ReleaseAlibabaAscpLogisticsCpGetAPIRequest 将 AlibabaAscpLogisticsCpGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsCpGetAPIRequest(v *AlibabaAscpLogisticsCpGetAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsCpGetAPIRequest.Put(v)
}
