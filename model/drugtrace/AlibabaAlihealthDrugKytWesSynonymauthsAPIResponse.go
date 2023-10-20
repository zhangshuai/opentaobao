package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesSynonymauthsAPIResponse 物流企业查询货主企业信息 API返回值
// alibaba.alihealth.drug.kyt.wes.synonymauths
//
// 物流企业查询货主企业信息
type AlibabaAlihealthDrugKytWesSynonymauthsAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesSynonymauthsAPIResponseModel
}

// AlibabaAlihealthDrugKytWesSynonymauthsAPIResponseModel is 物流企业查询货主企业信息 成功返回结果
type AlibabaAlihealthDrugKytWesSynonymauthsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_synonymauths_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytWesSynonymauthsResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
