package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品编辑 API返回值 
alibaba.idle.isv.item.edit

服务商ISV闲鱼商品编辑操作
*/
type AlibabaIdleIsvItemEditAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvItemEditResponse
}

// 服务商闲鱼商品编辑 成功返回结果
type AlibabaIdleIsvItemEditResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_item_edit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果result
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`
}