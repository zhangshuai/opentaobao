package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductskuaddAPIRequest 产品sku添加接口 API请求
// taobao.fenxiao.product.sku.add
//
// 添加产品SKU信息
type TaobaofenxiaoproductskuaddAPIRequest struct {
	model.Params
	// 采购基准价，最大值1000000000
	_standardPrice string
	// 代销采购价
	_agentCostPrice string
	// 经销采购价
	_dealerCostPrice string
	// 商家编码
	_skuNumber string
	// sku属性
	_properties string
	// 产品ID
	_productId int64
	// sku产品库存，在0到1000000之间，如果不传，则库存为0
	_quantity int64
}

// NewTaobaofenxiaoproductskuaddRequest 初始化TaobaofenxiaoproductskuaddAPIRequest对象
func NewTaobaofenxiaoproductskuaddRequest() *TaobaofenxiaoproductskuaddAPIRequest {
	return &TaobaofenxiaoproductskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductskuaddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandardPrice is StandardPrice Setter
// 采购基准价，最大值1000000000
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetStandardPrice(_standardPrice string) error {
	r._standardPrice = _standardPrice
	r.Set("standard_price", _standardPrice)
	return nil
}

// GetStandardPrice StandardPrice Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetStandardPrice() string {
	return r._standardPrice
}

// SetAgentCostPrice is AgentCostPrice Setter
// 代销采购价
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetAgentCostPrice(_agentCostPrice string) error {
	r._agentCostPrice = _agentCostPrice
	r.Set("agent_cost_price", _agentCostPrice)
	return nil
}

// GetAgentCostPrice AgentCostPrice Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetAgentCostPrice() string {
	return r._agentCostPrice
}

// SetDealerCostPrice is DealerCostPrice Setter
// 经销采购价
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetDealerCostPrice(_dealerCostPrice string) error {
	r._dealerCostPrice = _dealerCostPrice
	r.Set("dealer_cost_price", _dealerCostPrice)
	return nil
}

// GetDealerCostPrice DealerCostPrice Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetDealerCostPrice() string {
	return r._dealerCostPrice
}

// SetSkuNumber is SkuNumber Setter
// 商家编码
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetSkuNumber(_skuNumber string) error {
	r._skuNumber = _skuNumber
	r.Set("sku_number", _skuNumber)
	return nil
}

// GetSkuNumber SkuNumber Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetSkuNumber() string {
	return r._skuNumber
}

// SetProperties is Properties Setter
// sku属性
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetQuantity is Quantity Setter
// sku产品库存，在0到1000000之间，如果不传，则库存为0
func (r *TaobaofenxiaoproductskuaddAPIRequest) SetQuantity(_quantity int64) error {
	r._quantity = _quantity
	r.Set("quantity", _quantity)
	return nil
}

// GetQuantity Quantity Getter
func (r TaobaofenxiaoproductskuaddAPIRequest) GetQuantity() int64 {
	return r._quantity
}
