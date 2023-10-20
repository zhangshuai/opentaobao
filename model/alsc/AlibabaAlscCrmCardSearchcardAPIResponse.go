package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardsearchcardAPIResponse 搜索卡实例列表(支持号段查询) API返回值
// alibaba.alsc.crm.card.searchcard
//
// 搜索卡实例列表(支持号段查询)
type AlibabaalsccrmcardsearchcardAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmcardsearchcardAPIResponseModel
}

// AlibabaalsccrmcardsearchcardAPIResponseModel is 搜索卡实例列表(支持号段查询) 成功返回结果
type AlibabaalsccrmcardsearchcardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_searchcard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
