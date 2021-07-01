package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息通知 API请求
alibaba.legal.case.common.notice

同步通知给诉讼系统
*/
type AlibabaLegalCaseCommonNoticeAPIRequest struct {
    model.Params
    // 案件id
    _caseId   int64
    // 委托id
    _entrustId   int64
    // adminicular_evidence（补充证据）risk_early_warning（风险预警）
    _type   string
    // 消息模型
    _noticeModel   *NoticeModel
}

// 初始化AlibabaLegalCaseCommonNoticeAPIRequest对象
func NewAlibabaLegalCaseCommonNoticeRequest() *AlibabaLegalCaseCommonNoticeAPIRequest{
    return &AlibabaLegalCaseCommonNoticeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCommonNoticeAPIRequest) GetApiMethodName() string {
    return "alibaba.legal.case.common.notice"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCommonNoticeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CaseId Setter
// 案件id
func (r *AlibabaLegalCaseCommonNoticeAPIRequest) SetCaseId(_caseId int64) error {
    r._caseId = _caseId
    r.Set("case_id", _caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseCommonNoticeAPIRequest) GetCaseId() int64 {
    return r._caseId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseCommonNoticeAPIRequest) SetEntrustId(_entrustId int64) error {
    r._entrustId = _entrustId
    r.Set("entrust_id", _entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseCommonNoticeAPIRequest) GetEntrustId() int64 {
    return r._entrustId
}
// Type Setter
// adminicular_evidence（补充证据）risk_early_warning（风险预警）
func (r *AlibabaLegalCaseCommonNoticeAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaLegalCaseCommonNoticeAPIRequest) GetType() string {
    return r._type
}
// NoticeModel Setter
// 消息模型
func (r *AlibabaLegalCaseCommonNoticeAPIRequest) SetNoticeModel(_noticeModel *NoticeModel) error {
    r._noticeModel = _noticeModel
    r.Set("notice_model", _noticeModel)
    return nil
}

// NoticeModel Getter
func (r AlibabaLegalCaseCommonNoticeAPIRequest) GetNoticeModel() *NoticeModel {
    return r._noticeModel
}