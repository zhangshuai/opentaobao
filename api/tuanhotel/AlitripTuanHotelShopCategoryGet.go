package tuanhotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tuanhotel"
)

/* 
商家店铺类目查询 
alitrip.tuan.hotel.shop.category.get

查询商家店铺类目信息
*/
func AlitripTuanHotelShopCategoryGet(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelShopCategoryGetRequest, session string) (*tuanhotel.AlitripTuanHotelShopCategoryGetAPIResponse, error) {
    var resp tuanhotel.AlitripTuanHotelShopCategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}