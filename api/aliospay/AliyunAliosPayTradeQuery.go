package aliospay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayTradeQuery 查询支付结果接口
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
func AliyunAliosPayTradeQuery(ctx context.Context, clt *core.SDKClient, req *aliospay.AliyunAliosPayTradeQueryAPIRequest, resp *aliospay.AliyunAliosPayTradeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
