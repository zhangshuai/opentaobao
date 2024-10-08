package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoAutofinanceLoanReceive 接收放款结果通知
// tmall.aliauto.autofinance.loan.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
func TmallAliautoAutofinanceLoanReceive(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoAutofinanceLoanReceiveAPIRequest, resp *tmallcar.TmallAliautoAutofinanceLoanReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
