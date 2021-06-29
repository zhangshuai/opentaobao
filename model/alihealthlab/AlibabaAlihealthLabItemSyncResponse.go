package alihealthlab

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康检验检测商品发布 API返回值 
alibaba.alihealth.lab.item.sync

iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU
*/
type AlibabaAlihealthLabItemSyncAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthLabItemSyncResponse
}

// 阿里健康检验检测商品发布 成功返回结果
type AlibabaAlihealthLabItemSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_lab_item_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // SUCCESS - 成功，FAIL - 失败，UNKNOWN - 未知或参数异常
    ResultStatus   string `json:"result_status,omitempty" xml:"result_status,omitempty"`
    // 可读的结果码（错误码）
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}