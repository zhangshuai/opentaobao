package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugtableAPIResponse 查询药品目录信息 API返回值
// alibaba.alihealth.drug.kyt.drugtable
//
// 查询药品目录信息
type AlibabaAlihealthDrugKytDrugtableAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrugtableAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugtableAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrugtableAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrugtableAPIResponseModel is 查询药品目录信息 成功返回结果
type AlibabaAlihealthDrugKytDrugtableAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugtable_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrugtableResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugtableAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrugtableAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugtableAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrugtableAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrugtableAPIResponse
func GetAlibabaAlihealthDrugKytDrugtableAPIResponse() *AlibabaAlihealthDrugKytDrugtableAPIResponse {
	return poolAlibabaAlihealthDrugKytDrugtableAPIResponse.Get().(*AlibabaAlihealthDrugKytDrugtableAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrugtableAPIResponse 将 AlibabaAlihealthDrugKytDrugtableAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrugtableAPIResponse(v *AlibabaAlihealthDrugKytDrugtableAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrugtableAPIResponse.Put(v)
}
