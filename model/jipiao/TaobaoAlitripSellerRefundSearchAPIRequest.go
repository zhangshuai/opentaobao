package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundsearchAPIRequest 【机票代理商】退票申请单列表 API请求
// taobao.alitrip.seller.refund.search
//
// 查询退票申请单列表
type TaobaoalitripsellerrefundsearchAPIRequest struct {
	model.Params
	// 结束时间
	_endTime string
	// 开始时间
	_startTime string
	// 申请单状态（如果为空查询所有状态，1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
	_status int64
}

// NewTaobaoalitripsellerrefundsearchRequest 初始化TaobaoalitripsellerrefundsearchAPIRequest对象
func NewTaobaoalitripsellerrefundsearchRequest() *TaobaoalitripsellerrefundsearchAPIRequest {
	return &TaobaoalitripsellerrefundsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripsellerrefundsearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripsellerrefundsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripsellerrefundsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoalitripsellerrefundsearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoalitripsellerrefundsearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoalitripsellerrefundsearchAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoalitripsellerrefundsearchAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetStatus is Status Setter
// 申请单状态（如果为空查询所有状态，1初始 2接受 3确认 4失败 5买家处理 6卖家处理 7等待小二回填 8退款成功）
func (r *TaobaoalitripsellerrefundsearchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoalitripsellerrefundsearchAPIRequest) GetStatus() int64 {
	return r._status
}
