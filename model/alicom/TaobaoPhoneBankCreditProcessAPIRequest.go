package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaophonebankcreditprocessAPIRequest 虚拟话费任务银行信用卡办理进度回传 API请求
// taobao.phone.bank.credit.process
//
// 虚拟话费任务银行信用卡办理进度回传
type TaobaophonebankcreditprocessAPIRequest struct {
	model.Params
	// 办卡进度的回调
	_bankCreditRequest *BankCreditProcessRequest
}

// NewTaobaophonebankcreditprocessRequest 初始化TaobaophonebankcreditprocessAPIRequest对象
func NewTaobaophonebankcreditprocessRequest() *TaobaophonebankcreditprocessAPIRequest {
	return &TaobaophonebankcreditprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaophonebankcreditprocessAPIRequest) GetApiMethodName() string {
	return "taobao.phone.bank.credit.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaophonebankcreditprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaophonebankcreditprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBankCreditRequest is BankCreditRequest Setter
// 办卡进度的回调
func (r *TaobaophonebankcreditprocessAPIRequest) SetBankCreditRequest(_bankCreditRequest *BankCreditProcessRequest) error {
	r._bankCreditRequest = _bankCreditRequest
	r.Set("bank_credit_request", _bankCreditRequest)
	return nil
}

// GetBankCreditRequest BankCreditRequest Getter
func (r TaobaophonebankcreditprocessAPIRequest) GetBankCreditRequest() *BankCreditProcessRequest {
	return r._bankCreditRequest
}