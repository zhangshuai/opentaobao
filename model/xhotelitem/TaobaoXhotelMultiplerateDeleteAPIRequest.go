package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultiplerateDeleteAPIRequest 复杂价格删除 API请求
// taobao.xhotel.multiplerate.delete
//
// 酒店产品库rate删除
type TaobaoXhotelMultiplerateDeleteAPIRequest struct {
	model.Params
	// 渠道，和推送房价所使用的渠道保持一致
	_vendor string
	// 商家价格政策编码
	_rateplanCode string
	// 商家房型编码
	_outRid string
	// 连住天数
	_occupancy int64
	// 入住人数
	_lengthofstay int64
}

// NewTaobaoXhotelMultiplerateDeleteRequest 初始化TaobaoXhotelMultiplerateDeleteAPIRequest对象
func NewTaobaoXhotelMultiplerateDeleteRequest() *TaobaoXhotelMultiplerateDeleteAPIRequest {
	return &TaobaoXhotelMultiplerateDeleteAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelMultiplerateDeleteAPIRequest) Reset() {
	r._vendor = ""
	r._rateplanCode = ""
	r._outRid = ""
	r._occupancy = 0
	r._lengthofstay = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.multiplerate.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 渠道，和推送房价所使用的渠道保持一致
func (r *TaobaoXhotelMultiplerateDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelMultiplerateDeleteAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetOutRid is OutRid Setter
// 商家房型编码
func (r *TaobaoXhotelMultiplerateDeleteAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetOccupancy is Occupancy Setter
// 连住天数
func (r *TaobaoXhotelMultiplerateDeleteAPIRequest) SetOccupancy(_occupancy int64) error {
	r._occupancy = _occupancy
	r.Set("occupancy", _occupancy)
	return nil
}

// GetOccupancy Occupancy Getter
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetOccupancy() int64 {
	return r._occupancy
}

// SetLengthofstay is Lengthofstay Setter
// 入住人数
func (r *TaobaoXhotelMultiplerateDeleteAPIRequest) SetLengthofstay(_lengthofstay int64) error {
	r._lengthofstay = _lengthofstay
	r.Set("lengthofstay", _lengthofstay)
	return nil
}

// GetLengthofstay Lengthofstay Getter
func (r TaobaoXhotelMultiplerateDeleteAPIRequest) GetLengthofstay() int64 {
	return r._lengthofstay
}

var poolTaobaoXhotelMultiplerateDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelMultiplerateDeleteRequest()
	},
}

// GetTaobaoXhotelMultiplerateDeleteRequest 从 sync.Pool 获取 TaobaoXhotelMultiplerateDeleteAPIRequest
func GetTaobaoXhotelMultiplerateDeleteAPIRequest() *TaobaoXhotelMultiplerateDeleteAPIRequest {
	return poolTaobaoXhotelMultiplerateDeleteAPIRequest.Get().(*TaobaoXhotelMultiplerateDeleteAPIRequest)
}

// ReleaseTaobaoXhotelMultiplerateDeleteAPIRequest 将 TaobaoXhotelMultiplerateDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelMultiplerateDeleteAPIRequest(v *TaobaoXhotelMultiplerateDeleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelMultiplerateDeleteAPIRequest.Put(v)
}
