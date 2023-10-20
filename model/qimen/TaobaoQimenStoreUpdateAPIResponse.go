package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreUpdateAPIResponse 门店更新接口 API返回值
// taobao.qimen.store.update
//
// 商家在ERP等系统中调用该接口，更新门店信息
type TaobaoQimenStoreUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStoreUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStoreUpdateAPIResponseModel).Reset()
}

// TaobaoQimenStoreUpdateAPIResponseModel is 门店更新接口 成功返回结果
type TaobaoQimenStoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStoreUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTaobaoQimenStoreUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreUpdateAPIResponse)
	},
}

// GetTaobaoQimenStoreUpdateAPIResponse 从 sync.Pool 获取 TaobaoQimenStoreUpdateAPIResponse
func GetTaobaoQimenStoreUpdateAPIResponse() *TaobaoQimenStoreUpdateAPIResponse {
	return poolTaobaoQimenStoreUpdateAPIResponse.Get().(*TaobaoQimenStoreUpdateAPIResponse)
}

// ReleaseTaobaoQimenStoreUpdateAPIResponse 将 TaobaoQimenStoreUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStoreUpdateAPIResponse(v *TaobaoQimenStoreUpdateAPIResponse) {
	v.Reset()
	poolTaobaoQimenStoreUpdateAPIResponse.Put(v)
}
