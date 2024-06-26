package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterReservecondDeleteAPIRequest 删除主动预约开通条件 API请求
// tmall.servicecenter.reservecond.delete
//
// 删除主动预约开通条件
type TmallServicecenterReservecondDeleteAPIRequest struct {
	model.Params
	// 主动预约条件删除
	_reserveOpenConditionDelDto *ReserveOpenConditionDelDto
}

// NewTmallServicecenterReservecondDeleteRequest 初始化TmallServicecenterReservecondDeleteAPIRequest对象
func NewTmallServicecenterReservecondDeleteRequest() *TmallServicecenterReservecondDeleteAPIRequest {
	return &TmallServicecenterReservecondDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterReservecondDeleteAPIRequest) Reset() {
	r._reserveOpenConditionDelDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondDeleteAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterReservecondDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveOpenConditionDelDto is ReserveOpenConditionDelDto Setter
// 主动预约条件删除
func (r *TmallServicecenterReservecondDeleteAPIRequest) SetReserveOpenConditionDelDto(_reserveOpenConditionDelDto *ReserveOpenConditionDelDto) error {
	r._reserveOpenConditionDelDto = _reserveOpenConditionDelDto
	r.Set("reserve_open_condition_del_dto", _reserveOpenConditionDelDto)
	return nil
}

// GetReserveOpenConditionDelDto ReserveOpenConditionDelDto Getter
func (r TmallServicecenterReservecondDeleteAPIRequest) GetReserveOpenConditionDelDto() *ReserveOpenConditionDelDto {
	return r._reserveOpenConditionDelDto
}

var poolTmallServicecenterReservecondDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterReservecondDeleteRequest()
	},
}

// GetTmallServicecenterReservecondDeleteRequest 从 sync.Pool 获取 TmallServicecenterReservecondDeleteAPIRequest
func GetTmallServicecenterReservecondDeleteAPIRequest() *TmallServicecenterReservecondDeleteAPIRequest {
	return poolTmallServicecenterReservecondDeleteAPIRequest.Get().(*TmallServicecenterReservecondDeleteAPIRequest)
}

// ReleaseTmallServicecenterReservecondDeleteAPIRequest 将 TmallServicecenterReservecondDeleteAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterReservecondDeleteAPIRequest(v *TmallServicecenterReservecondDeleteAPIRequest) {
	v.Reset()
	poolTmallServicecenterReservecondDeleteAPIRequest.Put(v)
}
