package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新ISV库存 API返回值 
alibaba.health.nr.cep.warstqty.batchupdate

青岛医保服务-ISV批量更新孔雀翎中库存数据
*/
type AlibabaHealthNrCepWarstqtyBatchupdateAPIResponse struct {
    model.CommonResponse
    AlibabaHealthNrCepWarstqtyBatchupdateResponse
}

// 批量更新ISV库存 成功返回结果
type AlibabaHealthNrCepWarstqtyBatchupdateResponse struct {
    XMLName xml.Name `xml:"alibaba_health_nr_cep_warstqty_batchupdate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 服务出参
    Result   *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}
