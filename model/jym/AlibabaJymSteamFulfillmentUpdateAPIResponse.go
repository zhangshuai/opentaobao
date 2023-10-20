package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymsteamfulfillmentupdateAPIResponse 交易猫Steam类目发履约态变更 API返回值
// alibaba.jym.steam.fulfillment.update
//
// 交易猫Steam类目发履约态变更
type AlibabajymsteamfulfillmentupdateAPIResponse struct {
	model.CommonResponse
	AlibabajymsteamfulfillmentupdateAPIResponseModel
}

// AlibabajymsteamfulfillmentupdateAPIResponseModel is 交易猫Steam类目发履约态变更 成功返回结果
type AlibabajymsteamfulfillmentupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_steam_fulfillment_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子集错误码，提供与业务细节使用
	SubCodeType string `json:"sub_code_type,omitempty" xml:"sub_code_type,omitempty"`
	// 子集错误描述，提供与业务细节使用
	SubExtraErrMsg string `json:"sub_extra_err_msg,omitempty" xml:"sub_extra_err_msg,omitempty"`
	// 错误码枚举
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误详细描述
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 执行结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 调用结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}