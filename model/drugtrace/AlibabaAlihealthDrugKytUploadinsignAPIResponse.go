package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadinsignAPIResponse 仓库批量扫码回传接口 API返回值
// alibaba.alihealth.drug.kyt.uploadinsign
//
// 连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。
type AlibabaAlihealthDrugKytUploadinsignAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUploadinsignAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadinsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUploadinsignAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUploadinsignAPIResponseModel is 仓库批量扫码回传接口 成功返回结果
type AlibabaAlihealthDrugKytUploadinsignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadinsign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码(BILL_DECODE_ERROR 单据转码失败 BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功(true 成功 false 失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadinsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytUploadinsignAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUploadinsignAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUploadinsignAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUploadinsignAPIResponse
func GetAlibabaAlihealthDrugKytUploadinsignAPIResponse() *AlibabaAlihealthDrugKytUploadinsignAPIResponse {
	return poolAlibabaAlihealthDrugKytUploadinsignAPIResponse.Get().(*AlibabaAlihealthDrugKytUploadinsignAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUploadinsignAPIResponse 将 AlibabaAlihealthDrugKytUploadinsignAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUploadinsignAPIResponse(v *AlibabaAlihealthDrugKytUploadinsignAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUploadinsignAPIResponse.Put(v)
}
