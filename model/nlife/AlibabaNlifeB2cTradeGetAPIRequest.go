package nlife

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeGetAPIRequest 零售+平台查询订单 API请求
// alibaba.nlife.b2c.trade.get
//
// 查询零售+平台创建出来的订单详情
type AlibabaNlifeB2cTradeGetAPIRequest struct {
	model.Params
	// 零售+平台订单号,和out_trade_no不能同时为空
	_tradeNo string
	// 外部订单号，和trade_no不能同时为空
	_outTradeNo string
	// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
	_storeId string
}

// NewAlibabaNlifeB2cTradeGetRequest 初始化AlibabaNlifeB2cTradeGetAPIRequest对象
func NewAlibabaNlifeB2cTradeGetRequest() *AlibabaNlifeB2cTradeGetAPIRequest {
	return &AlibabaNlifeB2cTradeGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNlifeB2cTradeGetAPIRequest) Reset() {
	r._tradeNo = ""
	r._outTradeNo = ""
	r._storeId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNlifeB2cTradeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeNo is TradeNo Setter
// 零售+平台订单号,和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeGetAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaNlifeB2cTradeGetAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeGetAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabaNlifeB2cTradeGetAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}

// SetStoreId is StoreId Setter
// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
func (r *AlibabaNlifeB2cTradeGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaNlifeB2cTradeGetAPIRequest) GetStoreId() string {
	return r._storeId
}

var poolAlibabaNlifeB2cTradeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNlifeB2cTradeGetRequest()
	},
}

// GetAlibabaNlifeB2cTradeGetRequest 从 sync.Pool 获取 AlibabaNlifeB2cTradeGetAPIRequest
func GetAlibabaNlifeB2cTradeGetAPIRequest() *AlibabaNlifeB2cTradeGetAPIRequest {
	return poolAlibabaNlifeB2cTradeGetAPIRequest.Get().(*AlibabaNlifeB2cTradeGetAPIRequest)
}

// ReleaseAlibabaNlifeB2cTradeGetAPIRequest 将 AlibabaNlifeB2cTradeGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaNlifeB2cTradeGetAPIRequest(v *AlibabaNlifeB2cTradeGetAPIRequest) {
	v.Reset()
	poolAlibabaNlifeB2cTradeGetAPIRequest.Put(v)
}
