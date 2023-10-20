package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandoverpdfgetAPIRequest 获取面单PDF文件数据 API请求
// cainiao.global.handover.pdf.get
//
// 返回指定大包面单的PDF文件数据
type CainiaoglobalhandoverpdfgetAPIRequest struct {
	model.Params
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 用户信息
	_userInfo *UserInfoDto
	// 大包编号id
	_handoverContentId int64
	// 打印数据类型，1：面单、4：发货标签、512：交接清单
	_type int64
}

// NewCainiaoglobalhandoverpdfgetRequest 初始化CainiaoglobalhandoverpdfgetAPIRequest对象
func NewCainiaoglobalhandoverpdfgetRequest() *CainiaoglobalhandoverpdfgetAPIRequest {
	return &CainiaoglobalhandoverpdfgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.pdf.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoglobalhandoverpdfgetAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoglobalhandoverpdfgetAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoglobalhandoverpdfgetAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetHandoverContentId is HandoverContentId Setter
// 大包编号id
func (r *CainiaoglobalhandoverpdfgetAPIRequest) SetHandoverContentId(_handoverContentId int64) error {
	r._handoverContentId = _handoverContentId
	r.Set("handover_content_id", _handoverContentId)
	return nil
}

// GetHandoverContentId HandoverContentId Getter
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetHandoverContentId() int64 {
	return r._handoverContentId
}

// SetType is Type Setter
// 打印数据类型，1：面单、4：发货标签、512：交接清单
func (r *CainiaoglobalhandoverpdfgetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r CainiaoglobalhandoverpdfgetAPIRequest) GetType() int64 {
	return r._type
}
