package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenRechargeOperateAPIResponse 储值操作接口 API返回值
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
type AlibabaAlscCrmOpenRechargeOperateAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenRechargeOperateAPIResponseModel
}

// AlibabaAlscCrmOpenRechargeOperateAPIResponseModel is 储值操作接口 成功返回结果
type AlibabaAlscCrmOpenRechargeOperateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_recharge_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
