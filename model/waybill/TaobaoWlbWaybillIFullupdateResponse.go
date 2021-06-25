package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
面单信息更新接口v1.0 APIResponse
taobao.wlb.waybill.i.fullupdate

商家更新电子面单号对应的订单信息。<br/><br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；<br/><br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
*/
type TaobaoWlbWaybillIFullupdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWaybillIFullupdateResponse `json:"taobao_wlb_waybill_i_fullupdate_response,omitempty"`
}

type TaobaoWlbWaybillIFullupdateResponse struct {

    // 更新接口出参
    WaybillApplyUpdateInfo   *WaybillApplyUpdateInfo `json:"waybill_apply_update_info,omitempty"`

}