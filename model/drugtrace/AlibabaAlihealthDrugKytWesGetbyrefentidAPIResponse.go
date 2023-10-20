package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesgetbyrefentidAPIResponse 根据企业唯一标识查看企业详细信息 API返回值
// alibaba.alihealth.drug.kyt.wes.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
type AlibabaalihealthdrugkytwesgetbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytwesgetbyrefentidAPIResponseModel
}

// AlibabaalihealthdrugkytwesgetbyrefentidAPIResponseModel is 根据企业唯一标识查看企业详细信息 成功返回结果
type AlibabaalihealthdrugkytwesgetbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_getbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytwesgetbyrefentidResultModel `json:"result,omitempty" xml:"result,omitempty"`
}