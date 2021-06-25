package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
网点/门店状态修改 APIResponse
tmall.servicecenter.servicestore.updatestatus

修改网点/门店状态
*/
type TmallServicecenterServicestoreUpdatestatusAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServicestoreUpdatestatusResponse `json:"tmall_servicecenter_servicestore_updatestatus_response,omitempty"`
}

type TmallServicecenterServicestoreUpdatestatusResponse struct {

    // 方法调用结果
    Result   *TmallServicecenterServicestoreUpdatestatusResult `json:"result,omitempty"`

}