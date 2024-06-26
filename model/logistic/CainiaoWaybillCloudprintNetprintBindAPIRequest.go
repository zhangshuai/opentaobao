package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillCloudprintNetprintBindAPIRequest 网络打印机绑定 API请求
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
type CainiaoWaybillCloudprintNetprintBindAPIRequest struct {
	model.Params
	// req
	_params *CloudPrinterBindRequest
}

// NewCainiaoWaybillCloudprintNetprintBindRequest 初始化CainiaoWaybillCloudprintNetprintBindAPIRequest对象
func NewCainiaoWaybillCloudprintNetprintBindRequest() *CainiaoWaybillCloudprintNetprintBindAPIRequest {
	return &CainiaoWaybillCloudprintNetprintBindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillCloudprintNetprintBindAPIRequest) Reset() {
	r._params = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.cloudprint.netprint.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParams is Params Setter
// req
func (r *CainiaoWaybillCloudprintNetprintBindAPIRequest) SetParams(_params *CloudPrinterBindRequest) error {
	r._params = _params
	r.Set("params", _params)
	return nil
}

// GetParams Params Getter
func (r CainiaoWaybillCloudprintNetprintBindAPIRequest) GetParams() *CloudPrinterBindRequest {
	return r._params
}

var poolCainiaoWaybillCloudprintNetprintBindAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillCloudprintNetprintBindRequest()
	},
}

// GetCainiaoWaybillCloudprintNetprintBindRequest 从 sync.Pool 获取 CainiaoWaybillCloudprintNetprintBindAPIRequest
func GetCainiaoWaybillCloudprintNetprintBindAPIRequest() *CainiaoWaybillCloudprintNetprintBindAPIRequest {
	return poolCainiaoWaybillCloudprintNetprintBindAPIRequest.Get().(*CainiaoWaybillCloudprintNetprintBindAPIRequest)
}

// ReleaseCainiaoWaybillCloudprintNetprintBindAPIRequest 将 CainiaoWaybillCloudprintNetprintBindAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillCloudprintNetprintBindAPIRequest(v *CainiaoWaybillCloudprintNetprintBindAPIRequest) {
	v.Reset()
	poolCainiaoWaybillCloudprintNetprintBindAPIRequest.Put(v)
}
