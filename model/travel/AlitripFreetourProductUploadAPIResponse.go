package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFreetourProductUploadAPIResponse 自由行商品发布及编辑接口 API返回值
// alitrip.freetour.product.upload
//
// 自由行 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
type AlitripFreetourProductUploadAPIResponse struct {
	model.CommonResponse
	AlitripFreetourProductUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripFreetourProductUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripFreetourProductUploadAPIResponseModel).Reset()
}

// AlitripFreetourProductUploadAPIResponseModel is 自由行商品发布及编辑接口 成功返回结果
type AlitripFreetourProductUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_freetour_product_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// firstResult
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripFreetourProductUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = nil
}

var poolAlitripFreetourProductUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripFreetourProductUploadAPIResponse)
	},
}

// GetAlitripFreetourProductUploadAPIResponse 从 sync.Pool 获取 AlitripFreetourProductUploadAPIResponse
func GetAlitripFreetourProductUploadAPIResponse() *AlitripFreetourProductUploadAPIResponse {
	return poolAlitripFreetourProductUploadAPIResponse.Get().(*AlitripFreetourProductUploadAPIResponse)
}

// ReleaseAlitripFreetourProductUploadAPIResponse 将 AlitripFreetourProductUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripFreetourProductUploadAPIResponse(v *AlitripFreetourProductUploadAPIResponse) {
	v.Reset()
	poolAlitripFreetourProductUploadAPIResponse.Put(v)
}
