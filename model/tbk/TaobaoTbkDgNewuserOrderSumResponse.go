package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-拉新活动对应数据查询 APIResponse
taobao.tbk.dg.newuser.order.sum

拉新活动汇总API
*/
type TaobaoTbkDgNewuserOrderSumAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkDgNewuserOrderSumResponse `json:"taobao_tbk_dg_newuser_order_sum_response,omitempty"`
}

type TaobaoTbkDgNewuserOrderSumResponse struct {

    // data
    Results   *TaobaoTbkDgNewuserOrderSumData `json:"results,omitempty"`

}