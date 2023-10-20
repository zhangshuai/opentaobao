package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretaildevicepayUrlgetAPIRequest 贩卖机支付二维链接获取 API请求
// alibaba.retail.device.payUrl.get
//
// 贩卖机支付二维链接获取
type AlibabaretaildevicepayUrlgetAPIRequest struct {
	model.Params
	// 外部订单id
	_isvOrderId string
	// 业务名称
	_bizName string
	// 设备sn
	_deviceId string
	// 商品id
	_itemId int64
	// 1表示商品box，0或者为空表示普通商品
	_itemType int64
}

// NewAlibabaretaildevicepayUrlgetRequest 初始化AlibabaretaildevicepayUrlgetAPIRequest对象
func NewAlibabaretaildevicepayUrlgetRequest() *AlibabaretaildevicepayUrlgetAPIRequest {
	return &AlibabaretaildevicepayUrlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.payUrl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvOrderId is IsvOrderId Setter
// 外部订单id
func (r *AlibabaretaildevicepayUrlgetAPIRequest) SetIsvOrderId(_isvOrderId string) error {
	r._isvOrderId = _isvOrderId
	r.Set("isv_order_id", _isvOrderId)
	return nil
}

// GetIsvOrderId IsvOrderId Getter
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetIsvOrderId() string {
	return r._isvOrderId
}

// SetBizName is BizName Setter
// 业务名称
func (r *AlibabaretaildevicepayUrlgetAPIRequest) SetBizName(_bizName string) error {
	r._bizName = _bizName
	r.Set("biz_name", _bizName)
	return nil
}

// GetBizName BizName Getter
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetBizName() string {
	return r._bizName
}

// SetDeviceId is DeviceId Setter
// 设备sn
func (r *AlibabaretaildevicepayUrlgetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaretaildevicepayUrlgetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemType is ItemType Setter
// 1表示商品box，0或者为空表示普通商品
func (r *AlibabaretaildevicepayUrlgetAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r AlibabaretaildevicepayUrlgetAPIRequest) GetItemType() int64 {
	return r._itemType
}
