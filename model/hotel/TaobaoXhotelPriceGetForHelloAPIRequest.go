package hotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelPriceGetForHelloAPIRequest 哈罗合作方获取酒店库存报价 API请求
// taobao.xhotel.price.get.for.hello
//
// 哈罗合作项目，供哈罗合作方按需查询已开通城市下的标准酒店下指定时间段内的库存报价信息；在用户登录方面，返回结果不涉及用户个人信息，不涉及商家信息；仅根据不同用户，查询对应会员等级后，返回不同报价；
type TaobaoXhotelPriceGetForHelloAPIRequest struct {
	model.Params
	// 参数封装
	_hotelPriceParam *HotelPriceParam
}

// NewTaobaoXhotelPriceGetForHelloRequest 初始化TaobaoXhotelPriceGetForHelloAPIRequest对象
func NewTaobaoXhotelPriceGetForHelloRequest() *TaobaoXhotelPriceGetForHelloAPIRequest {
	return &TaobaoXhotelPriceGetForHelloAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelPriceGetForHelloAPIRequest) Reset() {
	r._hotelPriceParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelPriceGetForHelloAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.price.get.for.hello"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelPriceGetForHelloAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelPriceGetForHelloAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHotelPriceParam is HotelPriceParam Setter
// 参数封装
func (r *TaobaoXhotelPriceGetForHelloAPIRequest) SetHotelPriceParam(_hotelPriceParam *HotelPriceParam) error {
	r._hotelPriceParam = _hotelPriceParam
	r.Set("hotel_price_param", _hotelPriceParam)
	return nil
}

// GetHotelPriceParam HotelPriceParam Getter
func (r TaobaoXhotelPriceGetForHelloAPIRequest) GetHotelPriceParam() *HotelPriceParam {
	return r._hotelPriceParam
}

var poolTaobaoXhotelPriceGetForHelloAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelPriceGetForHelloRequest()
	},
}

// GetTaobaoXhotelPriceGetForHelloRequest 从 sync.Pool 获取 TaobaoXhotelPriceGetForHelloAPIRequest
func GetTaobaoXhotelPriceGetForHelloAPIRequest() *TaobaoXhotelPriceGetForHelloAPIRequest {
	return poolTaobaoXhotelPriceGetForHelloAPIRequest.Get().(*TaobaoXhotelPriceGetForHelloAPIRequest)
}

// ReleaseTaobaoXhotelPriceGetForHelloAPIRequest 将 TaobaoXhotelPriceGetForHelloAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelPriceGetForHelloAPIRequest(v *TaobaoXhotelPriceGetForHelloAPIRequest) {
	v.Reset()
	poolTaobaoXhotelPriceGetForHelloAPIRequest.Put(v)
}
