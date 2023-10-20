package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicedeductgetAPIRequest 发票扣减的接口 API请求
// alibaba.einvoice.deduct.get
//
// 获取历史发票扣减量、每日发票扣减量的接口
type AlibabaeinvoicedeductgetAPIRequest struct {
	model.Params
	// 税号
	_payeeRegisterNo string
	// 业务日期
	_bizDate string
	// 类型 1：所有 2：当日
	_type int64
}

// NewAlibabaeinvoicedeductgetRequest 初始化AlibabaeinvoicedeductgetAPIRequest对象
func NewAlibabaeinvoicedeductgetRequest() *AlibabaeinvoicedeductgetAPIRequest {
	return &AlibabaeinvoicedeductgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicedeductgetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.deduct.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicedeductgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicedeductgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税号
func (r *AlibabaeinvoicedeductgetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicedeductgetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetBizDate is BizDate Setter
// 业务日期
func (r *AlibabaeinvoicedeductgetAPIRequest) SetBizDate(_bizDate string) error {
	r._bizDate = _bizDate
	r.Set("biz_date", _bizDate)
	return nil
}

// GetBizDate BizDate Getter
func (r AlibabaeinvoicedeductgetAPIRequest) GetBizDate() string {
	return r._bizDate
}

// SetType is Type Setter
// 类型 1：所有 2：当日
func (r *AlibabaeinvoicedeductgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaeinvoicedeductgetAPIRequest) GetType() int64 {
	return r._type
}
