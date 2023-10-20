package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyUploadretailAPIResponse 医院上传出库信息 API返回值
// alibaba.alihealth.drug.kyt.yy.uploadretail
//
// 医院上传出库信息接口
type AlibabaAlihealthDrugKytYyUploadretailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytYyUploadretailAPIResponseModel
}

// AlibabaAlihealthDrugKytYyUploadretailAPIResponseModel is 医院上传出库信息 成功返回结果
type AlibabaAlihealthDrugKytYyUploadretailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_uploadretail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传单据文件队列表标识
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码(BILL_DECODE_ERROR 单据转码失败 2.BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作是否成功(true 成功 ,false失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
