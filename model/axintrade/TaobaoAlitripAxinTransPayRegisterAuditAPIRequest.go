package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterAuditAPIRequest 阿信支付入驻审核通知 API请求
// taobao.alitrip.axin.trans.pay.register.audit
//
// 阿信支付入驻审核通知
type TaobaoAlitripAxinTransPayRegisterAuditAPIRequest struct {
	model.Params
	// 支付入驻审核对象
	_axinPayRegisterAuditDto *AxinPayRegisterAuditDto
}

// NewTaobaoAlitripAxinTransPayRegisterAuditRequest 初始化TaobaoAlitripAxinTransPayRegisterAuditAPIRequest对象
func NewTaobaoAlitripAxinTransPayRegisterAuditRequest() *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest {
	return &TaobaoAlitripAxinTransPayRegisterAuditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) Reset() {
	r._axinPayRegisterAuditDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinPayRegisterAuditDto is AxinPayRegisterAuditDto Setter
// 支付入驻审核对象
func (r *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) SetAxinPayRegisterAuditDto(_axinPayRegisterAuditDto *AxinPayRegisterAuditDto) error {
	r._axinPayRegisterAuditDto = _axinPayRegisterAuditDto
	r.Set("axin_pay_register_audit_dto", _axinPayRegisterAuditDto)
	return nil
}

// GetAxinPayRegisterAuditDto AxinPayRegisterAuditDto Getter
func (r TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) GetAxinPayRegisterAuditDto() *AxinPayRegisterAuditDto {
	return r._axinPayRegisterAuditDto
}

var poolTaobaoAlitripAxinTransPayRegisterAuditAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripAxinTransPayRegisterAuditRequest()
	},
}

// GetTaobaoAlitripAxinTransPayRegisterAuditRequest 从 sync.Pool 获取 TaobaoAlitripAxinTransPayRegisterAuditAPIRequest
func GetTaobaoAlitripAxinTransPayRegisterAuditAPIRequest() *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest {
	return poolTaobaoAlitripAxinTransPayRegisterAuditAPIRequest.Get().(*TaobaoAlitripAxinTransPayRegisterAuditAPIRequest)
}

// ReleaseTaobaoAlitripAxinTransPayRegisterAuditAPIRequest 将 TaobaoAlitripAxinTransPayRegisterAuditAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripAxinTransPayRegisterAuditAPIRequest(v *TaobaoAlitripAxinTransPayRegisterAuditAPIRequest) {
	v.Reset()
	poolTaobaoAlitripAxinTransPayRegisterAuditAPIRequest.Put(v)
}
