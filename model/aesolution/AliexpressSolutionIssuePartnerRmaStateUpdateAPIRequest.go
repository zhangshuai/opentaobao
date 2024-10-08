package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest aliexpress.solution.issue.partner.rma.state.update API请求
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
type AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest struct {
	model.Params
	// RMA's order state update request
	_rmaStateUpdateRequest *RmaStateUpdateRequest
}

// NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest 初始化AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest对象
func NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest {
	return &AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) Reset() {
	r._rmaStateUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.state.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRmaStateUpdateRequest is RmaStateUpdateRequest Setter
// RMA&#39;s order state update request
func (r *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) SetRmaStateUpdateRequest(_rmaStateUpdateRequest *RmaStateUpdateRequest) error {
	r._rmaStateUpdateRequest = _rmaStateUpdateRequest
	r.Set("rma_state_update_request", _rmaStateUpdateRequest)
	return nil
}

// GetRmaStateUpdateRequest RmaStateUpdateRequest Getter
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetRmaStateUpdateRequest() *RmaStateUpdateRequest {
	return r._rmaStateUpdateRequest
}

var poolAliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest()
	},
}

// GetAliexpressSolutionIssuePartnerRmaStateUpdateRequest 从 sync.Pool 获取 AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest
func GetAliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest() *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest {
	return poolAliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest.Get().(*AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest)
}

// ReleaseAliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest 将 AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest(v *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) {
	v.Reset()
	poolAliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest.Put(v)
}
