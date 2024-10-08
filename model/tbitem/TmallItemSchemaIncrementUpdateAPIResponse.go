package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSchemaIncrementUpdateAPIResponse 天猫根据规则增量更新商品 API返回值
// tmall.item.schema.increment.update
//
// 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
type TmallItemSchemaIncrementUpdateAPIResponse struct {
	model.CommonResponse
	TmallItemSchemaIncrementUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSchemaIncrementUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSchemaIncrementUpdateAPIResponseModel).Reset()
}

// TmallItemSchemaIncrementUpdateAPIResponseModel is 天猫根据规则增量更新商品 成功返回结果
type TmallItemSchemaIncrementUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_schema_increment_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回商品发布结果
	UpdateItemResult string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`
	// 商品更新操作成功时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSchemaIncrementUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateItemResult = ""
	m.GmtModified = ""
}

var poolTmallItemSchemaIncrementUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSchemaIncrementUpdateAPIResponse)
	},
}

// GetTmallItemSchemaIncrementUpdateAPIResponse 从 sync.Pool 获取 TmallItemSchemaIncrementUpdateAPIResponse
func GetTmallItemSchemaIncrementUpdateAPIResponse() *TmallItemSchemaIncrementUpdateAPIResponse {
	return poolTmallItemSchemaIncrementUpdateAPIResponse.Get().(*TmallItemSchemaIncrementUpdateAPIResponse)
}

// ReleaseTmallItemSchemaIncrementUpdateAPIResponse 将 TmallItemSchemaIncrementUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSchemaIncrementUpdateAPIResponse(v *TmallItemSchemaIncrementUpdateAPIResponse) {
	v.Reset()
	poolTmallItemSchemaIncrementUpdateAPIResponse.Put(v)
}
