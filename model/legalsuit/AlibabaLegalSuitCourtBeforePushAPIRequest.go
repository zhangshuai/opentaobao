package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或保存庭前信息 API请求
alibaba.legal.suit.court.before.push

更新或者保存庭前信息
*/
type AlibabaLegalSuitCourtBeforePushAPIRequest struct {
    model.Params
    // 庭前信息
    _beforeCourtModel   *BeforeCourtModel
}

// 初始化AlibabaLegalSuitCourtBeforePushAPIRequest对象
func NewAlibabaLegalSuitCourtBeforePushRequest() *AlibabaLegalSuitCourtBeforePushAPIRequest{
    return &AlibabaLegalSuitCourtBeforePushAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.before.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BeforeCourtModel Setter
// 庭前信息
func (r *AlibabaLegalSuitCourtBeforePushAPIRequest) SetBeforeCourtModel(_beforeCourtModel *BeforeCourtModel) error {
    r._beforeCourtModel = _beforeCourtModel
    r.Set("before_court_model", _beforeCourtModel)
    return nil
}

// BeforeCourtModel Getter
func (r AlibabaLegalSuitCourtBeforePushAPIRequest) GetBeforeCourtModel() *BeforeCourtModel {
    return r._beforeCourtModel
}