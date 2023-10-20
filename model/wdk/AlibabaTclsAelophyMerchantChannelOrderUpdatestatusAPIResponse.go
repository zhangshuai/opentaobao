package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponse 翱象商家自有渠道 订单状态更新 API返回值
// alibaba.tcls.aelophy.merchant.channel.order.updatestatus
//
// 订单状态变更
type AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponseModel
}

// AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponseModel is 翱象商家自有渠道 订单状态更新 成功返回结果
type AlibabatclsaelophymerchantchannelorderupdatestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_merchant_channel_order_updatestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabatclsaelophymerchantchannelorderupdatestatusApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
