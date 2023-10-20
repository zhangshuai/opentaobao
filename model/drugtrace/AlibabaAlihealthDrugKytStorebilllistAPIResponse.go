package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytstorebilllistAPIResponse 零售端平台单据查询 API返回值
// alibaba.alihealth.drug.kyt.storebilllist
//
// 零售端平台单据查询
type AlibabaalihealthdrugkytstorebilllistAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytstorebilllistAPIResponseModel
}

// AlibabaalihealthdrugkytstorebilllistAPIResponseModel is 零售端平台单据查询 成功返回结果
type AlibabaalihealthdrugkytstorebilllistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_storebilllist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytstorebilllistResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
