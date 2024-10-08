package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenSupplychainTrainTrade 商旅火车票交易流水接口
// alitrip.btrip.open.supplychain.train.trade
//
// 商旅火车票交易流水接口
func AlitripBtripOpenSupplychainTrainTrade(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainTrainTradeAPIRequest, resp *btrip.AlitripBtripOpenSupplychainTrainTradeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
