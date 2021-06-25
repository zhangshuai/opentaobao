package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改运费模板 APIResponse
taobao.delivery.template.update

修改运费模板
*/
type TaobaoDeliveryTemplateUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDeliveryTemplateUpdateResponse `json:"taobao_delivery_template_update_response,omitempty"`
}

type TaobaoDeliveryTemplateUpdateResponse struct {

    // 表示修改是否成功
    Complete   bool `json:"complete,omitempty"`

}