package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelProductsGetAPIRequest 查询供应商的产品数据 API请求
// tmall.channel.products.get
//
// 查询供应商的产品数据。
//
// * 入参传入pids将优先查询，即只按这个条件查询。
// *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
// * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
// * 入参fields传入images将查询多图数据，不传只返回主图数据。
// * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
// * 查询结果按照产品发布时间倒序，即时间近的数据在前。
// * 传入channel 渠道，会只返回相应渠道的产品
type TmallChannelProductsGetAPIRequest struct {
	model.Params
	// top_query_product_d_o
	_topQueryProductDO *TopQueryProductDo
}

// NewTmallChannelProductsGetRequest 初始化TmallChannelProductsGetAPIRequest对象
func NewTmallChannelProductsGetRequest() *TmallChannelProductsGetAPIRequest {
	return &TmallChannelProductsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallChannelProductsGetAPIRequest) Reset() {
	r._topQueryProductDO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelProductsGetAPIRequest) GetApiMethodName() string {
	return "tmall.channel.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelProductsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallChannelProductsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopQueryProductDO is TopQueryProductDO Setter
// top_query_product_d_o
func (r *TmallChannelProductsGetAPIRequest) SetTopQueryProductDO(_topQueryProductDO *TopQueryProductDo) error {
	r._topQueryProductDO = _topQueryProductDO
	r.Set("top_query_product_d_o", _topQueryProductDO)
	return nil
}

// GetTopQueryProductDO TopQueryProductDO Getter
func (r TmallChannelProductsGetAPIRequest) GetTopQueryProductDO() *TopQueryProductDo {
	return r._topQueryProductDO
}

var poolTmallChannelProductsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallChannelProductsGetRequest()
	},
}

// GetTmallChannelProductsGetRequest 从 sync.Pool 获取 TmallChannelProductsGetAPIRequest
func GetTmallChannelProductsGetAPIRequest() *TmallChannelProductsGetAPIRequest {
	return poolTmallChannelProductsGetAPIRequest.Get().(*TmallChannelProductsGetAPIRequest)
}

// ReleaseTmallChannelProductsGetAPIRequest 将 TmallChannelProductsGetAPIRequest 放入 sync.Pool
func ReleaseTmallChannelProductsGetAPIRequest(v *TmallChannelProductsGetAPIRequest) {
	v.Reset()
	poolTmallChannelProductsGetAPIRequest.Put(v)
}
