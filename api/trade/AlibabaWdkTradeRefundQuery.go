package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabawdktraderefundquery 外部渠道查询退货订单详情接口
// alibaba.wdk.trade.refund.query
//
// 该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
func Alibabawdktraderefundquery(clt *core.SDKClient, req *trade.AlibabawdktraderefundqueryAPIRequest, session string) (*trade.AlibabawdktraderefundqueryAPIResponse, error) {
	var resp trade.AlibabawdktraderefundqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
