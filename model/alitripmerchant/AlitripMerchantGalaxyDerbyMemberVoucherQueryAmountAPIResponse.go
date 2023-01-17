package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse 查询用户拥有的臻享卡数量 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
type AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel is 查询用户拥有的臻享卡数量 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_query_amount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountResponse `json:"result,omitempty" xml:"result,omitempty"`
}
