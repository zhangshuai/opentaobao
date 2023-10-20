package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingexpirepromotiondeleteAPIResponse 短保优惠删除 API返回值
// alibaba.hm.marketing.expire.promotion.delete
//
// 短保优惠删除
type AlibabahmmarketingexpirepromotiondeleteAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingexpirepromotiondeleteAPIResponseModel
}

// AlibabahmmarketingexpirepromotiondeleteAPIResponseModel is 短保优惠删除 成功返回结果
type AlibabahmmarketingexpirepromotiondeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_expire_promotion_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Datas []ExpirePromotionResult `json:"datas,omitempty" xml:"datas>expire_promotion_result,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}