package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
AG物流签收状态写接口 APIResponse
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG
*/
type TaobaoNextoneLogisticsSignUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoNextoneLogisticsSignUpdateResponse `json:"taobao_nextone_logistics_sign_update_response,omitempty"`
}

type TaobaoNextoneLogisticsSignUpdateResponse struct {

    // 结果
    Result   *TaobaoNextoneLogisticsSignUpdateResult `json:"result,omitempty"`

}