package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbWaybillShengxianGet 商家获取生鲜电子面单号
// taobao.wlb.waybill.shengxian.get
//
// 商家通过交易订单号获取电子面单接口
func TaobaoWlbWaybillShengxianGet(clt *core.SDKClient, req *wlb.TaobaoWlbWaybillShengxianGetAPIRequest, resp *wlb.TaobaoWlbWaybillShengxianGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
