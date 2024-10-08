package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterServicestoreUpdateservicestorecapacity 更新网点容量
// tmall.servicecenter.servicestore.updateservicestorecapacity
//
// 更新网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要更新的网点容量不存在，会更新失败。
// 网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
func TmallServicecenterServicestoreUpdateservicestorecapacity(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterServicestoreUpdateservicestorecapacityAPIRequest, resp *tmallsc.TmallServicecenterServicestoreUpdateservicestorecapacityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
