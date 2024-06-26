package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityAddAPIResponse 创建通用单品优惠活动 API返回值
// taobao.promotionmisc.common.item.activity.add
//
// 创建通用单品优惠活动。
// 1、该接口只创建活动的基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
// 2、同一卖家下的活动数量限制为30个，超过限制需先调用taobao.promotionmisc.common.item.activity.delete接口删除无用的活动后才可再创建新的活动
type TaobaoPromotionmiscCommonItemActivityAddAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemActivityAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscCommonItemActivityAddAPIResponseModel).Reset()
}

// TaobaoPromotionmiscCommonItemActivityAddAPIResponseModel is 创建通用单品优惠活动 成功返回结果
type TaobaoPromotionmiscCommonItemActivityAddAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_activity_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 是否创建成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscCommonItemActivityAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ActivityId = 0
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscCommonItemActivityAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscCommonItemActivityAddAPIResponse)
	},
}

// GetTaobaoPromotionmiscCommonItemActivityAddAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemActivityAddAPIResponse
func GetTaobaoPromotionmiscCommonItemActivityAddAPIResponse() *TaobaoPromotionmiscCommonItemActivityAddAPIResponse {
	return poolTaobaoPromotionmiscCommonItemActivityAddAPIResponse.Get().(*TaobaoPromotionmiscCommonItemActivityAddAPIResponse)
}

// ReleaseTaobaoPromotionmiscCommonItemActivityAddAPIResponse 将 TaobaoPromotionmiscCommonItemActivityAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemActivityAddAPIResponse(v *TaobaoPromotionmiscCommonItemActivityAddAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemActivityAddAPIResponse.Put(v)
}
