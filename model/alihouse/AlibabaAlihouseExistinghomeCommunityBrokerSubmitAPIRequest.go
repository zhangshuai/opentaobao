package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest 提交小区专家 API请求
// alibaba.alihouse.existinghome.community.broker.submit
//
// 提交小区专家
type AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest struct {
	model.Params
	// 小区专家关系
	_relationDto *SubmitCommunityAgentRelationDto
}

// NewAlibabaAlihouseExistinghomeCommunityBrokerSubmitRequest 初始化AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomeCommunityBrokerSubmitRequest() *AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest {
	return &AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.community.broker.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationDto is RelationDto Setter
// 小区专家关系
func (r *AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest) SetRelationDto(_relationDto *SubmitCommunityAgentRelationDto) error {
	r._relationDto = _relationDto
	r.Set("relation_dto", _relationDto)
	return nil
}

// GetRelationDto RelationDto Getter
func (r AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest) GetRelationDto() *SubmitCommunityAgentRelationDto {
	return r._relationDto
}
