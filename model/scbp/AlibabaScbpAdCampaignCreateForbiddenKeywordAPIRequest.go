package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest 创建屏蔽词 API请求
// alibaba.scbp.ad.campaign.create.forbidden.keyword
//
// 创建屏蔽词
type AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto
}

// NewAlibabascbpadcampaigncreateforbiddenkeywordRequest 初始化AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest对象
func NewAlibabascbpadcampaigncreateforbiddenkeywordRequest() *AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest {
	return &AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.create.forbidden.keyword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenKeywordBatchOperation is ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) SetForbiddenKeywordBatchOperation(_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
	r._forbiddenKeywordBatchOperation = _forbiddenKeywordBatchOperation
	r.Set("forbidden_keyword_batch_operation", _forbiddenKeywordBatchOperation)
	return nil
}

// GetForbiddenKeywordBatchOperation ForbiddenKeywordBatchOperation Getter
func (r AlibabascbpadcampaigncreateforbiddenkeywordAPIRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
	return r._forbiddenKeywordBatchOperation
}
