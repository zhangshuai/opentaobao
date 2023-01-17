package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityDeliveryAddrConfirmAPIRequest 用户收件地址确认 API请求
// taobao.de.activity.delivery.addr.confirm
//
// 用户收件地址确认
type TaobaoDeActivityDeliveryAddrConfirmAPIRequest struct {
	model.Params
	// 加密流水号
	_serialNumber string
	// 地址Sign
	_addressSign string
}

// NewTaobaoDeActivityDeliveryAddrConfirmRequest 初始化TaobaoDeActivityDeliveryAddrConfirmAPIRequest对象
func NewTaobaoDeActivityDeliveryAddrConfirmRequest() *TaobaoDeActivityDeliveryAddrConfirmAPIRequest {
	return &TaobaoDeActivityDeliveryAddrConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDeActivityDeliveryAddrConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.delivery.addr.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDeActivityDeliveryAddrConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDeActivityDeliveryAddrConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSerialNumber is SerialNumber Setter
// 加密流水号
func (r *TaobaoDeActivityDeliveryAddrConfirmAPIRequest) SetSerialNumber(_serialNumber string) error {
	r._serialNumber = _serialNumber
	r.Set("serial_number", _serialNumber)
	return nil
}

// GetSerialNumber SerialNumber Getter
func (r TaobaoDeActivityDeliveryAddrConfirmAPIRequest) GetSerialNumber() string {
	return r._serialNumber
}

// SetAddressSign is AddressSign Setter
// 地址Sign
func (r *TaobaoDeActivityDeliveryAddrConfirmAPIRequest) SetAddressSign(_addressSign string) error {
	r._addressSign = _addressSign
	r.Set("address_sign", _addressSign)
	return nil
}

// GetAddressSign AddressSign Getter
func (r TaobaoDeActivityDeliveryAddrConfirmAPIRequest) GetAddressSign() string {
	return r._addressSign
}
