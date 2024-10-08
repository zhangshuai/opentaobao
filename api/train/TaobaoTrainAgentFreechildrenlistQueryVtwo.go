package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentFreechildrenlistQueryVtwo 免费儿童列表查询
// taobao.train.agent.freechildrenlist.query.vtwo
//
// 免费儿童列表查询
func TaobaoTrainAgentFreechildrenlistQueryVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest, resp *train.TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
