package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest 查询屏蔽词 API请求
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest 初始化AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest对象
func NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest() *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest {
	return &AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.forbidden.keyword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolAlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignFindForbiddenKeywordRequest()
	},
}

// GetAlibabaScbpAdCampaignFindForbiddenKeywordRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest
func GetAlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest() *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest {
	return poolAlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest.Get().(*AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest 将 AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest(v *AlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignFindForbiddenKeywordAPIRequest.Put(v)
}
