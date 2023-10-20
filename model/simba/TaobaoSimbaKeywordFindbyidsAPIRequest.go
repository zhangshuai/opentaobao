package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordfindbyidsAPIRequest （新）根据一堆关键词ids获取关键词 API请求
// taobao.simba.keyword.findbyids
//
// 根据一个关键词Id列表取得一组关键词
type TaobaosimbakeywordfindbyidsAPIRequest struct {
	model.Params
	// 关键词ids
	_bidwordIds []string
}

// NewTaobaosimbakeywordfindbyidsRequest 初始化TaobaosimbakeywordfindbyidsAPIRequest对象
func NewTaobaosimbakeywordfindbyidsRequest() *TaobaosimbakeywordfindbyidsAPIRequest {
	return &TaobaosimbakeywordfindbyidsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordfindbyidsAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.findbyids"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordfindbyidsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordfindbyidsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordIds is BidwordIds Setter
// 关键词ids
func (r *TaobaosimbakeywordfindbyidsAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaosimbakeywordfindbyidsAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}
