package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增网点容量 APIResponse
tmall.servicecenter.servicestore.createservicestorecapacity

新增网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点容量已存在，会新增失败。
网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
*/
type TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServicestoreCreateservicestorecapacityResponse `json:"tmall_servicecenter_servicestore_createservicestorecapacity_response,omitempty"`
}

type TmallServicecenterServicestoreCreateservicestorecapacityResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}