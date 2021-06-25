package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单核销 APIResponse
tmall.workcard.identify

工单核销，当工单完成以后，通过调用此接口核销
可以按照多维度核销工单，
电器预约安装按照工单维度核销，必选字段workcard_id,buyer_id,identify_code，可选字段attrs，通过扩展字段attrs回传机器码，格式{sn:'机器码'}
*/
type TmallWorkcardIdentifyAPIResponse struct {
    model.CommonResponse
    Response *TmallWorkcardIdentifyResponse `json:"tmall_workcard_identify_response,omitempty"`
}

type TmallWorkcardIdentifyResponse struct {

    // result
    Result   *TmallWorkcardIdentifyResult `json:"result,omitempty"`

}