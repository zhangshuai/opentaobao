package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanOperationAPIRequest 定向推广-计划开启/暂停/删除 API请求
// alibaba.scbp.target.ad.plan.operation
//
// 定向推广-计划开启/暂停/删除
type AlibabaScbpTargetAdPlanOperationAPIRequest struct {
	model.Params
	// TopP4pModifyQuickCampaignDTO
	_topP4pModifyQuickCampaignDTO *TopP4pModifyQuickCampaignDto
}

// NewAlibabaScbpTargetAdPlanOperationRequest 初始化AlibabaScbpTargetAdPlanOperationAPIRequest对象
func NewAlibabaScbpTargetAdPlanOperationRequest() *AlibabaScbpTargetAdPlanOperationAPIRequest {
	return &AlibabaScbpTargetAdPlanOperationAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpTargetAdPlanOperationAPIRequest) Reset() {
	r._topP4pModifyQuickCampaignDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanOperationAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanOperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpTargetAdPlanOperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopP4pModifyQuickCampaignDTO is TopP4pModifyQuickCampaignDTO Setter
// TopP4pModifyQuickCampaignDTO
func (r *AlibabaScbpTargetAdPlanOperationAPIRequest) SetTopP4pModifyQuickCampaignDTO(_topP4pModifyQuickCampaignDTO *TopP4pModifyQuickCampaignDto) error {
	r._topP4pModifyQuickCampaignDTO = _topP4pModifyQuickCampaignDTO
	r.Set("top_p4p_modify_quick_campaign_d_t_o", _topP4pModifyQuickCampaignDTO)
	return nil
}

// GetTopP4pModifyQuickCampaignDTO TopP4pModifyQuickCampaignDTO Getter
func (r AlibabaScbpTargetAdPlanOperationAPIRequest) GetTopP4pModifyQuickCampaignDTO() *TopP4pModifyQuickCampaignDto {
	return r._topP4pModifyQuickCampaignDTO
}

var poolAlibabaScbpTargetAdPlanOperationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpTargetAdPlanOperationRequest()
	},
}

// GetAlibabaScbpTargetAdPlanOperationRequest 从 sync.Pool 获取 AlibabaScbpTargetAdPlanOperationAPIRequest
func GetAlibabaScbpTargetAdPlanOperationAPIRequest() *AlibabaScbpTargetAdPlanOperationAPIRequest {
	return poolAlibabaScbpTargetAdPlanOperationAPIRequest.Get().(*AlibabaScbpTargetAdPlanOperationAPIRequest)
}

// ReleaseAlibabaScbpTargetAdPlanOperationAPIRequest 将 AlibabaScbpTargetAdPlanOperationAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanOperationAPIRequest(v *AlibabaScbpTargetAdPlanOperationAPIRequest) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanOperationAPIRequest.Put(v)
}
