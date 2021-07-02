package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelAdaptStoreGet 酒店团购套餐关联适用门店
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
func AlitripTuanHotelAdaptStoreGet(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelAdaptStoreGetAPIRequest, session string) (*tuanhotel.AlitripTuanHotelAdaptStoreGetAPIResponse, error) {
	var resp tuanhotel.AlitripTuanHotelAdaptStoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
