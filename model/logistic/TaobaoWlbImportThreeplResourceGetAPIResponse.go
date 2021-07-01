package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮获取资源列表 API返回值 
taobao.wlb.import.threepl.resource.get

获取3pl直邮的发货可用资源
*/
type TaobaoWlbImportThreeplResourceGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbImportThreeplResourceGetAPIResponseModel
}

// 3PL直邮获取资源列表 成功返回结果
type TaobaoWlbImportThreeplResourceGetAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_import_threepl_resource_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}