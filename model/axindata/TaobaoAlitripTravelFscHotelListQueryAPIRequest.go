package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfschotellistqueryAPIRequest 标准酒店信息查询-供销平台 API请求
// taobao.alitrip.travel.fsc.hotel.list.query
//
// 供销平台标准酒店信息列表查询
type TaobaoalitriptravelfschotellistqueryAPIRequest struct {
	model.Params
	// 上次查询最大shid
	_lastMaxShid int64
	// 城市代码
	_cityCode int64
	// 每次查询限制条数
	_limit int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoalitriptravelfschotellistqueryRequest 初始化TaobaoalitriptravelfschotellistqueryAPIRequest对象
func NewTaobaoalitriptravelfschotellistqueryRequest() *TaobaoalitriptravelfschotellistqueryAPIRequest {
	return &TaobaoalitriptravelfschotellistqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.fsc.hotel.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLastMaxShid is LastMaxShid Setter
// 上次查询最大shid
func (r *TaobaoalitriptravelfschotellistqueryAPIRequest) SetLastMaxShid(_lastMaxShid int64) error {
	r._lastMaxShid = _lastMaxShid
	r.Set("last_max_shid", _lastMaxShid)
	return nil
}

// GetLastMaxShid LastMaxShid Getter
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetLastMaxShid() int64 {
	return r._lastMaxShid
}

// SetCityCode is CityCode Setter
// 城市代码
func (r *TaobaoalitriptravelfschotellistqueryAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetLimit is Limit Setter
// 每次查询限制条数
func (r *TaobaoalitriptravelfschotellistqueryAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetLimit() int64 {
	return r._limit
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelfschotellistqueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelfschotellistqueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}