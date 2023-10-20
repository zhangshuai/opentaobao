package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemskugetAPIRequest 获取SKU API请求
// taobao.item.sku.get
//
// 获取sku_id所对应的sku数据
// sku_id对应的sku要属于传入的nick对应的卖家
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoitemskugetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
	_fields string
	// Sku的id。可以通过taobao.item.seller.get得到
	_skuId int64
	// 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
	_numIid int64
}

// NewTaobaoitemskugetRequest 初始化TaobaoitemskugetAPIRequest对象
func NewTaobaoitemskugetRequest() *TaobaoitemskugetAPIRequest {
	return &TaobaoitemskugetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemskugetAPIRequest) GetApiMethodName() string {
	return "taobao.item.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemskugetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemskugetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。
func (r *TaobaoitemskugetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoitemskugetAPIRequest) GetFields() string {
	return r._fields
}

// SetSkuId is SkuId Setter
// Sku的id。可以通过taobao.item.seller.get得到
func (r *TaobaoitemskugetAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoitemskugetAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetNumIid is NumIid Setter
// 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。
func (r *TaobaoitemskugetAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemskugetAPIRequest) GetNumIid() int64 {
	return r._numIid
}
