package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceModelsAPIRequest 获取品牌下设备列表 API请求
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
type YunosTvpubadminDeviceModelsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 牌照方
	_license int64
	// 品牌ID
	_brandId int64
}

// NewYunosTvpubadminDeviceModelsRequest 初始化YunosTvpubadminDeviceModelsAPIRequest对象
func NewYunosTvpubadminDeviceModelsRequest() *YunosTvpubadminDeviceModelsAPIRequest {
	return &YunosTvpubadminDeviceModelsAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceModelsAPIRequest) Reset() {
	r._terminalType = ""
	r._license = 0
	r._brandId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceModelsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.models"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceModelsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceModelsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTerminalType is TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceModelsAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// GetTerminalType TerminalType Getter
func (r YunosTvpubadminDeviceModelsAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceModelsAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceModelsAPIRequest) GetLicense() int64 {
	return r._license
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *YunosTvpubadminDeviceModelsAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r YunosTvpubadminDeviceModelsAPIRequest) GetBrandId() int64 {
	return r._brandId
}

var poolYunosTvpubadminDeviceModelsAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceModelsRequest()
	},
}

// GetYunosTvpubadminDeviceModelsRequest 从 sync.Pool 获取 YunosTvpubadminDeviceModelsAPIRequest
func GetYunosTvpubadminDeviceModelsAPIRequest() *YunosTvpubadminDeviceModelsAPIRequest {
	return poolYunosTvpubadminDeviceModelsAPIRequest.Get().(*YunosTvpubadminDeviceModelsAPIRequest)
}

// ReleaseYunosTvpubadminDeviceModelsAPIRequest 将 YunosTvpubadminDeviceModelsAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceModelsAPIRequest(v *YunosTvpubadminDeviceModelsAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceModelsAPIRequest.Put(v)
}
