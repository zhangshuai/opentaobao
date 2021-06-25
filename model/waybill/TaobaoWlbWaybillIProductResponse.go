package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家查询物流商产品类型接口 APIResponse
taobao.wlb.waybill.i.product

商家可以查询物流商的产品类型和服务能力。
*/
type TaobaoWlbWaybillIProductAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWaybillIProductResponse `json:"taobao_wlb_waybill_i_product_response,omitempty"`
}

type TaobaoWlbWaybillIProductResponse struct {

    // 产品类型返回
    ProductTypes   []WaybillProductType `json:"product_types,omitempty"`

}