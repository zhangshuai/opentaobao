package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardextrachargecreateAPIRequest 创建工单额外收费项 API请求
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
type TmallservicecenterworkcardextrachargecreateAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 额外收费项列表
	_extraChargeItemList *WorkcardExtraChargeCreateTuple
}

// NewTmallservicecenterworkcardextrachargecreateRequest 初始化TmallservicecenterworkcardextrachargecreateAPIRequest对象
func NewTmallservicecenterworkcardextrachargecreateRequest() *TmallservicecenterworkcardextrachargecreateAPIRequest {
	return &TmallservicecenterworkcardextrachargecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardextrachargecreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.extracharge.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardextrachargecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardextrachargecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardextrachargecreateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardextrachargecreateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetExtraChargeItemList is ExtraChargeItemList Setter
// 额外收费项列表
func (r *TmallservicecenterworkcardextrachargecreateAPIRequest) SetExtraChargeItemList(_extraChargeItemList *WorkcardExtraChargeCreateTuple) error {
	r._extraChargeItemList = _extraChargeItemList
	r.Set("extra_charge_item_list", _extraChargeItemList)
	return nil
}

// GetExtraChargeItemList ExtraChargeItemList Getter
func (r TmallservicecenterworkcardextrachargecreateAPIRequest) GetExtraChargeItemList() *WorkcardExtraChargeCreateTuple {
	return r._extraChargeItemList
}
