package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证码列表查询 APIResponse
taobao.vmarket.eticket.codes.get

查询某个订单的所有码的列表
*/
type TaobaoVmarketEticketCodesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketCodesGetResponse `json:"taobao_vmarket_eticket_codes_get_response,omitempty"`
}

type TaobaoVmarketEticketCodesGetResponse struct {

    // 记录总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 电子凭证码列表
    EticketCodes   []EticketCode `json:"eticket_codes,omitempty"`

}