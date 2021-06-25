package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人客诉数据上传 APIResponse
alibaba.dchain.miaoshifu.customer.complaints.put

数字服务供应链平台提供给服务商上传工人客诉数据
*/
type AlibabaDchainMiaoshifuCustomerComplaintsPutAPIResponse struct {
    model.CommonResponse
    Response *AlibabaDchainMiaoshifuCustomerComplaintsPutResponse `json:"alibaba_dchain_miaoshifu_customer_complaints_put_response,omitempty"`
}

type AlibabaDchainMiaoshifuCustomerComplaintsPutResponse struct {

    // 对外展示的错误信息
    DisplayMsg   string `json:"display_msg,omitempty"`

    // 是否成功，true：成功，false失败，当未false时，result_value为null
    ResultSuccess   bool `json:"result_success,omitempty"`

    // 错误码，异常返回码
    ResultErrorCode   string `json:"result_error_code,omitempty"`

    // 客诉记录唯一标识ID
    ResultValue   int64 `json:"result_value,omitempty"`

    // 错误信息
    ResultErrorMsg   string `json:"result_error_msg,omitempty"`

}