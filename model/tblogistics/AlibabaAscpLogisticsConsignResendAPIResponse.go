package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsConsignResendAPIResponse 修改物流公司和运单号 API返回值
// alibaba.ascp.logistics.consign.resend
//
// 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：
// 1、必须是已发货订单，自己联系发货的必须50天内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
// 2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司
type AlibabaAscpLogisticsConsignResendAPIResponse struct {
	model.CommonResponse
	AlibabaAscpLogisticsConsignResendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsConsignResendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpLogisticsConsignResendAPIResponseModel).Reset()
}

// AlibabaAscpLogisticsConsignResendAPIResponseModel is 修改物流公司和运单号 成功返回结果
type AlibabaAscpLogisticsConsignResendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_consign_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaAscpLogisticsConsignResendResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpLogisticsConsignResendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpLogisticsConsignResendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsConsignResendAPIResponse)
	},
}

// GetAlibabaAscpLogisticsConsignResendAPIResponse 从 sync.Pool 获取 AlibabaAscpLogisticsConsignResendAPIResponse
func GetAlibabaAscpLogisticsConsignResendAPIResponse() *AlibabaAscpLogisticsConsignResendAPIResponse {
	return poolAlibabaAscpLogisticsConsignResendAPIResponse.Get().(*AlibabaAscpLogisticsConsignResendAPIResponse)
}

// ReleaseAlibabaAscpLogisticsConsignResendAPIResponse 将 AlibabaAscpLogisticsConsignResendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpLogisticsConsignResendAPIResponse(v *AlibabaAscpLogisticsConsignResendAPIResponse) {
	v.Reset()
	poolAlibabaAscpLogisticsConsignResendAPIResponse.Put(v)
}
