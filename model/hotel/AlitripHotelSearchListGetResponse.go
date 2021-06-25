package hotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
酒店搜索List接口 APIResponse
alitrip.hotel.search.list.get

酒店搜索List接口
*/
type AlitripHotelSearchListGetAPIResponse struct {
    model.CommonResponse
    Response *AlitripHotelSearchListGetResponse `json:"alitrip_hotel_search_list_get_response,omitempty"`
}

type AlitripHotelSearchListGetResponse struct {

    // result
    Result   *AlitripHotelSearchListGetResult `json:"result,omitempty"`

}