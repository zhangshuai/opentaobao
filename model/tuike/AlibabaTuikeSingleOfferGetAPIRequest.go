package tuike

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询1688推客平台卖家推广中的商品信息 API请求
alibaba.tuike.single.offer.get

查询单个推客商品信息的接口
*/
type AlibabaTuikeSingleOfferGetAPIRequest struct {
    model.Params
    // 推客id
    _loginId   string
    // 商品id
    _offerId   int64
}

// 初始化AlibabaTuikeSingleOfferGetAPIRequest对象
func NewAlibabaTuikeSingleOfferGetRequest() *AlibabaTuikeSingleOfferGetAPIRequest{
    return &AlibabaTuikeSingleOfferGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTuikeSingleOfferGetAPIRequest) GetApiMethodName() string {
    return "alibaba.tuike.single.offer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTuikeSingleOfferGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LoginId Setter
// 推客id
func (r *AlibabaTuikeSingleOfferGetAPIRequest) SetLoginId(_loginId string) error {
    r._loginId = _loginId
    r.Set("login_id", _loginId)
    return nil
}

// LoginId Getter
func (r AlibabaTuikeSingleOfferGetAPIRequest) GetLoginId() string {
    return r._loginId
}
// OfferId Setter
// 商品id
func (r *AlibabaTuikeSingleOfferGetAPIRequest) SetOfferId(_offerId int64) error {
    r._offerId = _offerId
    r.Set("offer_id", _offerId)
    return nil
}

// OfferId Getter
func (r AlibabaTuikeSingleOfferGetAPIRequest) GetOfferId() int64 {
    return r._offerId
}