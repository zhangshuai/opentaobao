package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest 查询上游出库单明细(带追溯码信息) API请求
// alibaba.alihealth.drug.bill.upbill.detail.withcode
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 单据编码
	_billCode string
	// 发货企业renEntId
	_fromRefUserId string
	// 收货企业refEntId
	_toRefUserId string
	// 委托企业id
	_agentRefEntId string
}

// NewAlibabaalihealthdrugbillupbilldetailwithcodeRequest 初始化AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest对象
func NewAlibabaalihealthdrugbillupbilldetailwithcodeRequest() *AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest {
	return &AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.bill.upbill.detail.withcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetFromRefUserId is FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// GetFromRefUserId FromRefUserId Getter
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// SetToRefUserId is ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// GetToRefUserId ToRefUserId Getter
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}

// SetAgentRefEntId is AgentRefEntId Setter
// 委托企业id
func (r *AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugbillupbilldetailwithcodeAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}
