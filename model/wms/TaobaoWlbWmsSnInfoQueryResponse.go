package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单据序列号信息 API返回值 
taobao.wlb.wms.sn.info.query

查询仓库作业的各类单据记录的Sn信息
*/
type TaobaoWlbWmsSnInfoQueryAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsSnInfoQueryResponse
}

// 查询单据序列号信息 成功返回结果
type TaobaoWlbWmsSnInfoQueryResponse struct {
    XMLName xml.Name `xml:"wlb_wms_sn_info_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回
    Result   *TaobaoWlbWmsSnInfoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
