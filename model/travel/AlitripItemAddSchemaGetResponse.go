package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布模板 APIResponse
alitrip.item.add.schema.get

发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemAddSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *AlitripItemAddSchemaGetResponse `json:"alitrip_item_add_schema_get_response,omitempty"`
}

type AlitripItemAddSchemaGetResponse struct {

    // schema模板数据
    SchemaXmlFields   string `json:"schema_xml_fields,omitempty"`

}