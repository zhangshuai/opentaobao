package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopTrainExceedapplyGet 商旅火车票第三方超标审批单搜索接口
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
func AlitripBtripCorpopTrainExceedapplyGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopTrainExceedapplyGetAPIRequest, resp *btrip.AlitripBtripCorpopTrainExceedapplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
