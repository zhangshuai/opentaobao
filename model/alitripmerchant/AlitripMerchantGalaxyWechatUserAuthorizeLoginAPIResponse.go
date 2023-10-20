package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatuserauthorizeloginAPIResponse DFC-ID用户手机号授权登录 API返回值
// alitrip.merchant.galaxy.wechat.user.authorize.login
//
// DFC-ID用户手机号授权登录
type AlitripmerchantgalaxywechatuserauthorizeloginAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxywechatuserauthorizeloginAPIResponseModel
}

// AlitripmerchantgalaxywechatuserauthorizeloginAPIResponseModel is DFC-ID用户手机号授权登录 成功返回结果
type AlitripmerchantgalaxywechatuserauthorizeloginAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_user_authorize_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripmerchantgalaxywechatuserauthorizeloginResponse `json:"result,omitempty" xml:"result,omitempty"`
}