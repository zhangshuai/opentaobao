package cainiaoncwl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoNcwlJhorderQueryAPIRequest 农村物流集货单查询接口 API请求
// cainiao.ncwl.jhorder.query
//
// 提供给接入商家，查询农村物流集货单
type CainiaoNcwlJhorderQueryAPIRequest struct {
	model.Params
	// 1
	_param0 *JhRequest
}

// NewCainiaoNcwlJhorderQueryRequest 初始化CainiaoNcwlJhorderQueryAPIRequest对象
func NewCainiaoNcwlJhorderQueryRequest() *CainiaoNcwlJhorderQueryAPIRequest {
	return &CainiaoNcwlJhorderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoNcwlJhorderQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.ncwl.jhorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoNcwlJhorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoNcwlJhorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 1
func (r *CainiaoNcwlJhorderQueryAPIRequest) SetParam0(_param0 *JhRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r CainiaoNcwlJhorderQueryAPIRequest) GetParam0() *JhRequest {
	return r._param0
}
