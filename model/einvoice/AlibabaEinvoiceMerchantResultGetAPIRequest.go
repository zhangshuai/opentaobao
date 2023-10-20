package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicemerchantresultgetAPIRequest 商家自研ERP开票结果获取 API请求
// alibaba.einvoice.merchant.result.get
//
// 商家自研ERP开票结果获取
type AlibabaeinvoicemerchantresultgetAPIRequest struct {
	model.Params
	// 收款方税务登记证号
	_payeeRegisterNo string
	// 电商平台代码。淘宝：taobao，天猫：tmall
	_platformCode string
	// 电商平台对应的订单号
	_platformTid string
	// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
	_serialNo string
}

// NewAlibabaeinvoicemerchantresultgetRequest 初始化AlibabaeinvoicemerchantresultgetAPIRequest对象
func NewAlibabaeinvoicemerchantresultgetRequest() *AlibabaeinvoicemerchantresultgetAPIRequest {
	return &AlibabaeinvoicemerchantresultgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.merchant.result.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 收款方税务登记证号
func (r *AlibabaeinvoicemerchantresultgetAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetPlatformCode is PlatformCode Setter
// 电商平台代码。淘宝：taobao，天猫：tmall
func (r *AlibabaeinvoicemerchantresultgetAPIRequest) SetPlatformCode(_platformCode string) error {
	r._platformCode = _platformCode
	r.Set("platform_code", _platformCode)
	return nil
}

// GetPlatformCode PlatformCode Getter
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetPlatformCode() string {
	return r._platformCode
}

// SetPlatformTid is PlatformTid Setter
// 电商平台对应的订单号
func (r *AlibabaeinvoicemerchantresultgetAPIRequest) SetPlatformTid(_platformTid string) error {
	r._platformTid = _platformTid
	r.Set("platform_tid", _platformTid)
	return nil
}

// GetPlatformTid PlatformTid Getter
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetPlatformTid() string {
	return r._platformTid
}

// SetSerialNo is SerialNo Setter
// 流水号 (serial_no)和(platform_code,platform_tid)必须填写其中一组,serial_no优先级更高
func (r *AlibabaeinvoicemerchantresultgetAPIRequest) SetSerialNo(_serialNo string) error {
	r._serialNo = _serialNo
	r.Set("serial_no", _serialNo)
	return nil
}

// GetSerialNo SerialNo Getter
func (r AlibabaeinvoicemerchantresultgetAPIRequest) GetSerialNo() string {
	return r._serialNo
}
