package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountCheckmobileAPIRequest 阿里体育会员系统--手机号验证接口 API请求
// alibaba.alisports.passport.account.checkmobile
//
// 验证三方用户的手机号，获取对应的阿里体育会员id
type AlibabaAlisportsPassportAccountCheckmobileAPIRequest struct {
	model.Params
	// 业务appkey
	_alispAppKey string
	// 调用时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 合作方用户ID
	_appUid string
	// 用户呢称
	_nick string
	// 手机号
	_mobile string
	// 性别 0未设置 1男 2女 3保密
	_gender string
	// 生日
	_birthday string
}

// NewAlibabaAlisportsPassportAccountCheckmobileRequest 初始化AlibabaAlisportsPassportAccountCheckmobileAPIRequest对象
func NewAlibabaAlisportsPassportAccountCheckmobileRequest() *AlibabaAlisportsPassportAccountCheckmobileAPIRequest {
	return &AlibabaAlisportsPassportAccountCheckmobileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.checkmobile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlispAppKey Setter
// 业务appkey
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// Get AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// Set is AlispTime Setter
// 调用时间戳
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// Get AlispTime Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// Set is AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// Get AlispSign Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// Set is AppUid Setter
// 合作方用户ID
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// Get AppUid Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAppUid() string {
	return r._appUid
}

// Set is Nick Setter
// 用户呢称
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetNick() string {
	return r._nick
}

// Set is Mobile Setter
// 手机号
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is Gender Setter
// 性别 0未设置 1男 2女 3保密
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// Get Gender Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetGender() string {
	return r._gender
}

// Set is Birthday Setter
// 生日
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// Get Birthday Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetBirthday() string {
	return r._birthday
}
