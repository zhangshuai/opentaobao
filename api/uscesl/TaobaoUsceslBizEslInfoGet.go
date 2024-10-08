package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizEslInfoGet 价签设备信息查询接口
// taobao.uscesl.biz.esl.info.get
//
// 价签设备信息查询接口
func TaobaoUsceslBizEslInfoGet(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizEslInfoGetAPIRequest, resp *uscesl.TaobaoUsceslBizEslInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
