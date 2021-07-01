package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫商品尺码表模板列表 API返回值 
tmall.item.sizemapping.templates.list

获取所有尺码表模板列表。
*/
type TmallItemSizemappingTemplatesListAPIResponse struct {
    model.CommonResponse
    TmallItemSizemappingTemplatesListAPIResponseModel
}

// 获取天猫商品尺码表模板列表 成功返回结果
type TmallItemSizemappingTemplatesListAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_item_sizemapping_templates_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 尺码表模板列表
    SizeMappingTemplates   []SizeMappingTemplate `json:"size_mapping_templates,omitempty" xml:"size_mapping_templates>size_mapping_template,omitempty"`
}