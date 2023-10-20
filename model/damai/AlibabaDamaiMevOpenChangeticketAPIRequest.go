package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenchangeticketAPIRequest 大麦换验平台-第三方对外开放-票单接口changeTicket API请求
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
type AlibabadamaimevopenchangeticketAPIRequest struct {
	model.Params
	// 入参ticketIdOpenParam
	_ticketIdOpenParam *TicketIdOpenParam
}

// NewAlibabadamaimevopenchangeticketRequest 初始化AlibabadamaimevopenchangeticketAPIRequest对象
func NewAlibabadamaimevopenchangeticketRequest() *AlibabadamaimevopenchangeticketAPIRequest {
	return &AlibabadamaimevopenchangeticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenchangeticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.changeticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenchangeticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenchangeticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketIdOpenParam is TicketIdOpenParam Setter
// 入参ticketIdOpenParam
func (r *AlibabadamaimevopenchangeticketAPIRequest) SetTicketIdOpenParam(_ticketIdOpenParam *TicketIdOpenParam) error {
	r._ticketIdOpenParam = _ticketIdOpenParam
	r.Set("ticket_id_open_param", _ticketIdOpenParam)
	return nil
}

// GetTicketIdOpenParam TicketIdOpenParam Getter
func (r AlibabadamaimevopenchangeticketAPIRequest) GetTicketIdOpenParam() *TicketIdOpenParam {
	return r._ticketIdOpenParam
}
