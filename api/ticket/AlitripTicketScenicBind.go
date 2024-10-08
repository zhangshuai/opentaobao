package ticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// AlitripTicketScenicBind 【门票API2.0】门票景点绑定接口
// alitrip.ticket.scenic.bind
//
// 门票景点绑定接口，用于建立阿里标准景点id与商家系统景点id的映射关系。该接口同时支持新建和修改映射关系，当用户没有为ali_scenic_id建立过映射关系时，则判断为新建映射关系，否则为修改。可以通过设置update_out_scenic_id来修改ali_scenic_id与out_scenic_id的映射关系。
func AlitripTicketScenicBind(ctx context.Context, clt *core.SDKClient, req *ticket.AlitripTicketScenicBindAPIRequest, resp *ticket.AlitripTicketScenicBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
