package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoFilesGet 业务文件获取
// taobao.files.get
//
// 获取业务方暂存给ISV的文件列表
func TaobaoFilesGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoFilesGetAPIRequest, resp *util.TaobaoFilesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
