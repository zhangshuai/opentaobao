package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenPointOperate 积分操作接口
// alibaba.alsc.crm.open.point.operate
//
// 同步积分接口
func AlibabaAlscCrmOpenPointOperate(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenPointOperateAPIRequest, resp *alsc.AlibabaAlscCrmOpenPointOperateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
