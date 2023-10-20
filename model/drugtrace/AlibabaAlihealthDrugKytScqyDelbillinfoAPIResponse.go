package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse 根据单据号删除单据 API返回值
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
type AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel
}

// AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel is 根据单据号删除单据 成功返回结果
type AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_delbillinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyDelbillinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
