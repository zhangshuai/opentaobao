package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemPromotionRuleGetAPIResponse 获取商品已生效营销活动更新规则 API返回值
// taobao.item.promotion.rule.get
//
// 获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
type TaobaoItemPromotionRuleGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemPromotionRuleGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemPromotionRuleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemPromotionRuleGetAPIResponseModel).Reset()
}

// TaobaoItemPromotionRuleGetAPIResponseModel is 获取商品已生效营销活动更新规则 成功返回结果
type TaobaoItemPromotionRuleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_promotion_rule_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品规则信息
	Rules []ItemPromotionRule `json:"rules,omitempty" xml:"rules>item_promotion_rule,omitempty"`
	// 商品是否命中更新规则
	Effec bool `json:"effec,omitempty" xml:"effec,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemPromotionRuleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rules = m.Rules[:0]
	m.Effec = false
}

var poolTaobaoItemPromotionRuleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemPromotionRuleGetAPIResponse)
	},
}

// GetTaobaoItemPromotionRuleGetAPIResponse 从 sync.Pool 获取 TaobaoItemPromotionRuleGetAPIResponse
func GetTaobaoItemPromotionRuleGetAPIResponse() *TaobaoItemPromotionRuleGetAPIResponse {
	return poolTaobaoItemPromotionRuleGetAPIResponse.Get().(*TaobaoItemPromotionRuleGetAPIResponse)
}

// ReleaseTaobaoItemPromotionRuleGetAPIResponse 将 TaobaoItemPromotionRuleGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemPromotionRuleGetAPIResponse(v *TaobaoItemPromotionRuleGetAPIResponse) {
	v.Reset()
	poolTaobaoItemPromotionRuleGetAPIResponse.Put(v)
}
