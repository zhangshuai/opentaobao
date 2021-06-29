package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单个订阅关系 API返回值 
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationQueryAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanItemSubscribeRelationQueryResponse
}

// 查询单个订阅关系 成功返回结果
type TaobaoBaichuanItemSubscribeRelationQueryResponse struct {
    XMLName xml.Name `xml:"baichuan_item_subscribe_relation_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoBaichuanItemSubscribeRelationQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
