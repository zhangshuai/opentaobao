package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmisccommonitemdetaildeleteAPIRequest 删除通用单品优惠详情 API请求
// taobao.promotionmisc.common.item.detail.delete
//
// 删除通用单品优惠详情。
type TaobaopromotionmisccommonitemdetaildeleteAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 优惠详情ID
	_detailId int64
}

// NewTaobaopromotionmisccommonitemdetaildeleteRequest 初始化TaobaopromotionmisccommonitemdetaildeleteAPIRequest对象
func NewTaobaopromotionmisccommonitemdetaildeleteRequest() *TaobaopromotionmisccommonitemdetaildeleteAPIRequest {
	return &TaobaopromotionmisccommonitemdetaildeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmisccommonitemdetaildeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmisccommonitemdetaildeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmisccommonitemdetaildeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaopromotionmisccommonitemdetaildeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmisccommonitemdetaildeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetDetailId is DetailId Setter
// 优惠详情ID
func (r *TaobaopromotionmisccommonitemdetaildeleteAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaopromotionmisccommonitemdetaildeleteAPIRequest) GetDetailId() int64 {
	return r._detailId
}
