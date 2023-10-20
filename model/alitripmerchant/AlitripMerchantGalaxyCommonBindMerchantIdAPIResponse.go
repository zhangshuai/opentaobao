package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycommonbindmerchantidAPIResponse 绑定用户和merchantID API返回值
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
type AlitripmerchantgalaxycommonbindmerchantidAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxycommonbindmerchantidAPIResponseModel
}

// AlitripmerchantgalaxycommonbindmerchantidAPIResponseModel is 绑定用户和merchantID 成功返回结果
type AlitripmerchantgalaxycommonbindmerchantidAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_common_bind_merchant_id_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripmerchantgalaxycommonbindmerchantidResponse `json:"result,omitempty" xml:"result,omitempty"`
}