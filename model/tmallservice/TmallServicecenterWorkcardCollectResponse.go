package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单揽件 APIResponse
tmall.servicecenter.workcard.collect

服务工单揽件接口
*/
type TmallServicecenterWorkcardCollectAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardCollectResponse `json:"tmall_servicecenter_workcard_collect_response,omitempty"`
}

type TmallServicecenterWorkcardCollectResponse struct {

    // 响应结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}