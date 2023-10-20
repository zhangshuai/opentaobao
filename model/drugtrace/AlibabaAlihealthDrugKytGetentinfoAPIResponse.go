package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytgetentinfoAPIResponse 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API返回值
// alibaba.alihealth.drug.kyt.getentinfo
//
// 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
type AlibabaalihealthdrugkytgetentinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytgetentinfoAPIResponseModel
}

// AlibabaalihealthdrugkytgetentinfoAPIResponseModel is 根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 成功返回结果
type AlibabaalihealthdrugkytgetentinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getentinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytgetentinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
