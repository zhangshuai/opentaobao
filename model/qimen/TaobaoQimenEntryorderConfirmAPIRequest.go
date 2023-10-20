package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderConfirmAPIRequest 入库单确认接口 API请求
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmAPIRequest struct {
	model.Params
	//
	_request *EntryOrderConfirmRequest
}

// NewTaobaoQimenEntryorderConfirmRequest 初始化TaobaoQimenEntryorderConfirmAPIRequest对象
func NewTaobaoQimenEntryorderConfirmRequest() *TaobaoQimenEntryorderConfirmAPIRequest {
	return &TaobaoQimenEntryorderConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.entryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenEntryorderConfirmAPIRequest) SetRequest(_request *EntryOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenEntryorderConfirmAPIRequest) GetRequest() *EntryOrderConfirmRequest {
	return r._request
}
