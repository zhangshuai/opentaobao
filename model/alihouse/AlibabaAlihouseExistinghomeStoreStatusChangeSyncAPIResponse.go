package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse 门店状态变更 API返回值
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
type AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponseModel is 门店状态变更 成功返回结果
type AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_status_change_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// aaa
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// aaa
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// aaa
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// aaa
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.Data = false
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse
func GetAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse() *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse 将 AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse(v *AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse.Put(v)
}
