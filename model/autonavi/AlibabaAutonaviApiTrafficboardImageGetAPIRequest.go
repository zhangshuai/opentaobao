package autonavi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAutonaviApiTrafficboardImageGetAPIRequest 交通看板-栅格情报获取 API请求
// alibaba.autonavi.api.trafficboard.image.get
//
// 获取指定情报板ID的二进制数据（图片）
type AlibabaAutonaviApiTrafficboardImageGetAPIRequest struct {
	model.Params
	// 设备id,  按照userid 的配置，决定是否需要
	_deviceid string
	// 批次,终端批次，按照userid 的配置，决定是否需要
	_batch string
	// 图片 id
	_panelid string
	// 图像尺寸（可选）,默认尺寸为原始大小(960x600) 参数为:width  x   height   (例如:960x600),参数不正确时返回原始大小
	_size string
	// 是否为宽高等比例（可选）,参数值 true（默认）,表示宽高等比例缩放 false:  按请求尺寸缩放
	_whscale string
	// 城市编码
	_adcodes string
}

// NewAlibabaAutonaviApiTrafficboardImageGetRequest 初始化AlibabaAutonaviApiTrafficboardImageGetAPIRequest对象
func NewAlibabaAutonaviApiTrafficboardImageGetRequest() *AlibabaAutonaviApiTrafficboardImageGetAPIRequest {
	return &AlibabaAutonaviApiTrafficboardImageGetAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) Reset() {
	r._deviceid = ""
	r._batch = ""
	r._panelid = ""
	r._size = ""
	r._whscale = ""
	r._adcodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetApiMethodName() string {
	return "alibaba.autonavi.api.trafficboard.image.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceid is Deviceid Setter
// 设备id,  按照userid 的配置，决定是否需要
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) SetDeviceid(_deviceid string) error {
	r._deviceid = _deviceid
	r.Set("deviceid", _deviceid)
	return nil
}

// GetDeviceid Deviceid Getter
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetDeviceid() string {
	return r._deviceid
}

// SetBatch is Batch Setter
// 批次,终端批次，按照userid 的配置，决定是否需要
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) SetBatch(_batch string) error {
	r._batch = _batch
	r.Set("batch", _batch)
	return nil
}

// GetBatch Batch Getter
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetBatch() string {
	return r._batch
}

// SetPanelid is Panelid Setter
// 图片 id
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) SetPanelid(_panelid string) error {
	r._panelid = _panelid
	r.Set("panelid", _panelid)
	return nil
}

// GetPanelid Panelid Getter
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetPanelid() string {
	return r._panelid
}

// SetSize is Size Setter
// 图像尺寸（可选）,默认尺寸为原始大小(960x600) 参数为:width  x   height   (例如:960x600),参数不正确时返回原始大小
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) SetSize(_size string) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetSize() string {
	return r._size
}

// SetWhscale is Whscale Setter
// 是否为宽高等比例（可选）,参数值 true（默认）,表示宽高等比例缩放 false:  按请求尺寸缩放
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) SetWhscale(_whscale string) error {
	r._whscale = _whscale
	r.Set("whscale", _whscale)
	return nil
}

// GetWhscale Whscale Getter
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetWhscale() string {
	return r._whscale
}

// SetAdcodes is Adcodes Setter
// 城市编码
func (r *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) SetAdcodes(_adcodes string) error {
	r._adcodes = _adcodes
	r.Set("adcodes", _adcodes)
	return nil
}

// GetAdcodes Adcodes Getter
func (r AlibabaAutonaviApiTrafficboardImageGetAPIRequest) GetAdcodes() string {
	return r._adcodes
}

var poolAlibabaAutonaviApiTrafficboardImageGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAutonaviApiTrafficboardImageGetRequest()
	},
}

// GetAlibabaAutonaviApiTrafficboardImageGetRequest 从 sync.Pool 获取 AlibabaAutonaviApiTrafficboardImageGetAPIRequest
func GetAlibabaAutonaviApiTrafficboardImageGetAPIRequest() *AlibabaAutonaviApiTrafficboardImageGetAPIRequest {
	return poolAlibabaAutonaviApiTrafficboardImageGetAPIRequest.Get().(*AlibabaAutonaviApiTrafficboardImageGetAPIRequest)
}

// ReleaseAlibabaAutonaviApiTrafficboardImageGetAPIRequest 将 AlibabaAutonaviApiTrafficboardImageGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAutonaviApiTrafficboardImageGetAPIRequest(v *AlibabaAutonaviApiTrafficboardImageGetAPIRequest) {
	v.Reset()
	poolAlibabaAutonaviApiTrafficboardImageGetAPIRequest.Put(v)
}
