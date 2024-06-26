package ioti

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItAlbumDeviceSendimageAPIRequest 相框设备厂测刷图接口 API请求
// alibaba.it.album.device.sendimage
//
// 提供传入电子相框设备mac，mac需属于厂测白名单设备，将设备刷新为系统默认的厂测图片
type AlibabaItAlbumDeviceSendimageAPIRequest struct {
	model.Params
	// 下发图片mac地址
	_mac string
}

// NewAlibabaItAlbumDeviceSendimageRequest 初始化AlibabaItAlbumDeviceSendimageAPIRequest对象
func NewAlibabaItAlbumDeviceSendimageRequest() *AlibabaItAlbumDeviceSendimageAPIRequest {
	return &AlibabaItAlbumDeviceSendimageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItAlbumDeviceSendimageAPIRequest) Reset() {
	r._mac = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetApiMethodName() string {
	return "alibaba.it.album.device.sendimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMac is Mac Setter
// 下发图片mac地址
func (r *AlibabaItAlbumDeviceSendimageAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaItAlbumDeviceSendimageAPIRequest) GetMac() string {
	return r._mac
}

var poolAlibabaItAlbumDeviceSendimageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItAlbumDeviceSendimageRequest()
	},
}

// GetAlibabaItAlbumDeviceSendimageRequest 从 sync.Pool 获取 AlibabaItAlbumDeviceSendimageAPIRequest
func GetAlibabaItAlbumDeviceSendimageAPIRequest() *AlibabaItAlbumDeviceSendimageAPIRequest {
	return poolAlibabaItAlbumDeviceSendimageAPIRequest.Get().(*AlibabaItAlbumDeviceSendimageAPIRequest)
}

// ReleaseAlibabaItAlbumDeviceSendimageAPIRequest 将 AlibabaItAlbumDeviceSendimageAPIRequest 放入 sync.Pool
func ReleaseAlibabaItAlbumDeviceSendimageAPIRequest(v *AlibabaItAlbumDeviceSendimageAPIRequest) {
	v.Reset()
	poolAlibabaItAlbumDeviceSendimageAPIRequest.Put(v)
}
