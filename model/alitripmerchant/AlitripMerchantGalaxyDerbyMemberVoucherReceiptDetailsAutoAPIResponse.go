package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponse v5.0付费会员卡开票抬头自匹配 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.auto
//
// v5.0付费会员卡开票抬头自匹配
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponseModel is v5.0付费会员卡开票抬头自匹配 成功返回结果
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_receipt_details_auto_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 抬头自匹配接口返回结果
	Result *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoResponse `json:"result,omitempty" xml:"result,omitempty"`
}