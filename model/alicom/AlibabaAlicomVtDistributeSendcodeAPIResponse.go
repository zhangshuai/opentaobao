package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalicomvtdistributesendcodeAPIResponse 通信业务外放发送验证码 API返回值
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
type AlibabaalicomvtdistributesendcodeAPIResponse struct {
	model.CommonResponse
	AlibabaalicomvtdistributesendcodeAPIResponseModel
}

// AlibabaalicomvtdistributesendcodeAPIResponseModel is 通信业务外放发送验证码 成功返回结果
type AlibabaalicomvtdistributesendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_vt_distribute_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}