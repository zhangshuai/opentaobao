package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审批单 API请求
alitrip.btrip.apply.get

获取单个审批单的详情数据
*/
type AlitripBtripApplyGetAPIRequest struct {
    model.Params
    // 外部审批单id
    _thirdpartApplyId   string
    // 阿里商旅审批单id
    _applyId   int64
    // 企业id
    _corpId   string
    // 审批单展示id
    _applyShowId   string
}

// 初始化AlitripBtripApplyGetAPIRequest对象
func NewAlitripBtripApplyGetRequest() *AlitripBtripApplyGetAPIRequest{
    return &AlitripBtripApplyGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApplyGetAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.apply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApplyGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ThirdpartApplyId Setter
// 外部审批单id
func (r *AlitripBtripApplyGetAPIRequest) SetThirdpartApplyId(_thirdpartApplyId string) error {
    r._thirdpartApplyId = _thirdpartApplyId
    r.Set("thirdpart_apply_id", _thirdpartApplyId)
    return nil
}

// ThirdpartApplyId Getter
func (r AlitripBtripApplyGetAPIRequest) GetThirdpartApplyId() string {
    return r._thirdpartApplyId
}
// ApplyId Setter
// 阿里商旅审批单id
func (r *AlitripBtripApplyGetAPIRequest) SetApplyId(_applyId int64) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlitripBtripApplyGetAPIRequest) GetApplyId() int64 {
    return r._applyId
}
// CorpId Setter
// 企业id
func (r *AlitripBtripApplyGetAPIRequest) SetCorpId(_corpId string) error {
    r._corpId = _corpId
    r.Set("corp_id", _corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripApplyGetAPIRequest) GetCorpId() string {
    return r._corpId
}
// ApplyShowId Setter
// 审批单展示id
func (r *AlitripBtripApplyGetAPIRequest) SetApplyShowId(_applyShowId string) error {
    r._applyShowId = _applyShowId
    r.Set("apply_show_id", _applyShowId)
    return nil
}

// ApplyShowId Getter
func (r AlitripBtripApplyGetAPIRequest) GetApplyShowId() string {
    return r._applyShowId
}