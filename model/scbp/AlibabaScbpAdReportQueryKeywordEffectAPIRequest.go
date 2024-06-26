package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportQueryKeywordEffectAPIRequest 关键词报告 API请求
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
type AlibabaScbpAdReportQueryKeywordEffectAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_keywordReportOperation *KeywordReportOperationDto
}

// NewAlibabaScbpAdReportQueryKeywordEffectRequest 初始化AlibabaScbpAdReportQueryKeywordEffectAPIRequest对象
func NewAlibabaScbpAdReportQueryKeywordEffectRequest() *AlibabaScbpAdReportQueryKeywordEffectAPIRequest {
	return &AlibabaScbpAdReportQueryKeywordEffectAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) Reset() {
	r._topContext = nil
	r._keywordReportOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.query.keyword.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetKeywordReportOperation is KeywordReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
	r._keywordReportOperation = _keywordReportOperation
	r.Set("keyword_report_operation", _keywordReportOperation)
	return nil
}

// GetKeywordReportOperation KeywordReportOperation Getter
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
	return r._keywordReportOperation
}

var poolAlibabaScbpAdReportQueryKeywordEffectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdReportQueryKeywordEffectRequest()
	},
}

// GetAlibabaScbpAdReportQueryKeywordEffectRequest 从 sync.Pool 获取 AlibabaScbpAdReportQueryKeywordEffectAPIRequest
func GetAlibabaScbpAdReportQueryKeywordEffectAPIRequest() *AlibabaScbpAdReportQueryKeywordEffectAPIRequest {
	return poolAlibabaScbpAdReportQueryKeywordEffectAPIRequest.Get().(*AlibabaScbpAdReportQueryKeywordEffectAPIRequest)
}

// ReleaseAlibabaScbpAdReportQueryKeywordEffectAPIRequest 将 AlibabaScbpAdReportQueryKeywordEffectAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdReportQueryKeywordEffectAPIRequest(v *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdReportQueryKeywordEffectAPIRequest.Put(v)
}
