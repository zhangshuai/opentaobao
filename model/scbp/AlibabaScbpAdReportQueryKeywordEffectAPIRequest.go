package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报告 API请求
alibaba.scbp.ad.report.query.keyword.effect

关键词报告
*/
type AlibabaScbpAdReportQueryKeywordEffectAPIRequest struct {
    model.Params
    // 请求参数
    _keywordReportOperation   *KeywordReportOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdReportQueryKeywordEffectAPIRequest对象
func NewAlibabaScbpAdReportQueryKeywordEffectRequest() *AlibabaScbpAdReportQueryKeywordEffectAPIRequest{
    return &AlibabaScbpAdReportQueryKeywordEffectAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.keyword.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
    r._keywordReportOperation = _keywordReportOperation
    r.Set("keyword_report_operation", _keywordReportOperation)
    return nil
}

// KeywordReportOperation Getter
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
    return r._keywordReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}