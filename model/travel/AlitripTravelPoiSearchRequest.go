package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POI信息查询 API请求
alitrip.travel.poi.search

POI信息查询，用于商品更新使用
*/
type AlitripTravelPoiSearchRequest struct {
    model.Params
    // POI信息ID；
    id   int64
    // POI信息名称
    name   string
    // POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他
    type   int64
}

// 初始化AlitripTravelPoiSearchRequest对象
func NewAlitripTravelPoiSearchRequest() *AlitripTravelPoiSearchRequest{
    return &AlitripTravelPoiSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelPoiSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.poi.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelPoiSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// POI信息ID；
func (r *AlitripTravelPoiSearchRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r AlitripTravelPoiSearchRequest) GetId() int64 {
    return r.id
}
// Name Setter
// POI信息名称
func (r *AlitripTravelPoiSearchRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlitripTravelPoiSearchRequest) GetName() string {
    return r.name
}
// Type Setter
// POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他
func (r *AlitripTravelPoiSearchRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlitripTravelPoiSearchRequest) GetType() int64 {
    return r.type
}
