package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosGoodsSetprice 价格变更接口
// alibaba.mos.goods.setprice
//
// 价格变更接口，供供应商修改价格时使用。
func AlibabaMosGoodsSetprice(clt *core.SDKClient, req *moscm.AlibabaMosGoodsSetpriceAPIRequest, resp *moscm.AlibabaMosGoodsSetpriceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
