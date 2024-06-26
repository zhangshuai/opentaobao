package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugWxinfoUploadAPIRequest 小程序数据回传 API请求
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
type AlibabaAlihealthDrugWxinfoUploadAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo string
	// 店铺名称
	_shopInfo string
	// 售货员信息
	_salerInfo string
	// 渠道
	_isvChannel string
}

// NewAlibabaAlihealthDrugWxinfoUploadRequest 初始化AlibabaAlihealthDrugWxinfoUploadAPIRequest对象
func NewAlibabaAlihealthDrugWxinfoUploadRequest() *AlibabaAlihealthDrugWxinfoUploadAPIRequest {
	return &AlibabaAlihealthDrugWxinfoUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) Reset() {
	r._userInfo = ""
	r._shopInfo = ""
	r._salerInfo = ""
	r._isvChannel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.wxinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetUserInfo(_userInfo string) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetUserInfo() string {
	return r._userInfo
}

// SetShopInfo is ShopInfo Setter
// 店铺名称
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetShopInfo(_shopInfo string) error {
	r._shopInfo = _shopInfo
	r.Set("shop_info", _shopInfo)
	return nil
}

// GetShopInfo ShopInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetShopInfo() string {
	return r._shopInfo
}

// SetSalerInfo is SalerInfo Setter
// 售货员信息
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetSalerInfo(_salerInfo string) error {
	r._salerInfo = _salerInfo
	r.Set("saler_info", _salerInfo)
	return nil
}

// GetSalerInfo SalerInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetSalerInfo() string {
	return r._salerInfo
}

// SetIsvChannel is IsvChannel Setter
// 渠道
func (r *AlibabaAlihealthDrugWxinfoUploadAPIRequest) SetIsvChannel(_isvChannel string) error {
	r._isvChannel = _isvChannel
	r.Set("isv_channel", _isvChannel)
	return nil
}

// GetIsvChannel IsvChannel Getter
func (r AlibabaAlihealthDrugWxinfoUploadAPIRequest) GetIsvChannel() string {
	return r._isvChannel
}

var poolAlibabaAlihealthDrugWxinfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugWxinfoUploadRequest()
	},
}

// GetAlibabaAlihealthDrugWxinfoUploadRequest 从 sync.Pool 获取 AlibabaAlihealthDrugWxinfoUploadAPIRequest
func GetAlibabaAlihealthDrugWxinfoUploadAPIRequest() *AlibabaAlihealthDrugWxinfoUploadAPIRequest {
	return poolAlibabaAlihealthDrugWxinfoUploadAPIRequest.Get().(*AlibabaAlihealthDrugWxinfoUploadAPIRequest)
}

// ReleaseAlibabaAlihealthDrugWxinfoUploadAPIRequest 将 AlibabaAlihealthDrugWxinfoUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugWxinfoUploadAPIRequest(v *AlibabaAlihealthDrugWxinfoUploadAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugWxinfoUploadAPIRequest.Put(v)
}
