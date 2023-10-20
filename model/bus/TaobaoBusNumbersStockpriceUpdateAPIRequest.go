package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusnumbersstockpriceupdateAPIRequest 汽车票更新价格库存 API请求
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
type TaobaobusnumbersstockpriceupdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq
}

// NewTaobaobusnumbersstockpriceupdateRequest 初始化TaobaobusnumbersstockpriceupdateAPIRequest对象
func NewTaobaobusnumbersstockpriceupdateRequest() *TaobaobusnumbersstockpriceupdateAPIRequest {
	return &TaobaobusnumbersstockpriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusnumbersstockpriceupdateAPIRequest) GetApiMethodName() string {
	return "taobao.bus.numbers.stockprice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusnumbersstockpriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusnumbersstockpriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopBusPriceAndStockUpdateRQ is ParamTopBusPriceAndStockUpdateRQ Setter
// 请求参数
func (r *TaobaobusnumbersstockpriceupdateAPIRequest) SetParamTopBusPriceAndStockUpdateRQ(_paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq) error {
	r._paramTopBusPriceAndStockUpdateRQ = _paramTopBusPriceAndStockUpdateRQ
	r.Set("param_top_bus_price_and_stock_update_r_q", _paramTopBusPriceAndStockUpdateRQ)
	return nil
}

// GetParamTopBusPriceAndStockUpdateRQ ParamTopBusPriceAndStockUpdateRQ Getter
func (r TaobaobusnumbersstockpriceupdateAPIRequest) GetParamTopBusPriceAndStockUpdateRQ() *TopBusPriceAndStockUpdateRq {
	return r._paramTopBusPriceAndStockUpdateRQ
}
