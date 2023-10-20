package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoCarrierdriverLocation 查询骑手当前位置
// alibaba.ele.fengniao.carrierdriver.location
//
// 查询骑手当前位置
func AlibabaEleFengniaoCarrierdriverLocation(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCarrierdriverLocationAPIRequest, resp *logistic.AlibabaEleFengniaoCarrierdriverLocationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
