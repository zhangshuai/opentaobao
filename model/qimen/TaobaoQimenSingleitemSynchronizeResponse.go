package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 API返回值 
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS
*/
type TaobaoQimenSingleitemSynchronizeAPIResponse struct {
    model.CommonResponse
    TaobaoQimenSingleitemSynchronizeResponse
}

// 商品同步接口 成功返回结果
type TaobaoQimenSingleitemSynchronizeResponse struct {
    XMLName xml.Name `xml:"qimen_singleitem_synchronize_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
