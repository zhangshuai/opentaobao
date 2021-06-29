package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询追溯码对应的药品信息（疫情） API返回值 
alibaba.alihealth.drug.code.kyt.yq.querycode

通过追溯码码得到 药品名称、包装规格、剂型、剂型规格”、有效期至等信息。
*/
type AlibabaAlihealthDrugCodeKytYqQuerycodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugCodeKytYqQuerycodeResponse
}

// 查询追溯码对应的药品信息（疫情） 成功返回结果
type AlibabaAlihealthDrugCodeKytYqQuerycodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_yq_querycode_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 最外层结果
    Result   *AlibabaAlihealthDrugCodeKytYqQuerycodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
