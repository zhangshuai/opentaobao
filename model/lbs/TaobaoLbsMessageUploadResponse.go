package lbs

import (
    "github.com/bububa/opentaobao/model"
)

/* 
lbs数据采集 APIResponse
taobao.lbs.message.upload

lbs数据采集
*/
type TaobaoLbsMessageUploadAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLbsMessageUploadResponse `json:"taobao_lbs_message_upload_response,omitempty"`
}

type TaobaoLbsMessageUploadResponse struct {

    // result
    Result   bool `json:"result,omitempty"`

    // subCode
    ResultMsg   string `json:"result_msg,omitempty"`

    // subMsg
    ResultCode   string `json:"result_code,omitempty"`

}