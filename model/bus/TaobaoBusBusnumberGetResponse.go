package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票车次查询 APIResponse
taobao.bus.busnumber.get

提供汽车票车次查询服务
*/
type TaobaoBusBusnumberGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusBusnumberGetResponse `json:"taobao_bus_busnumber_get_response,omitempty"`
}

type TaobaoBusBusnumberGetResponse struct {

    // result
    Result   *TaobaoBusBusnumberGetResultSet `json:"result,omitempty"`

}