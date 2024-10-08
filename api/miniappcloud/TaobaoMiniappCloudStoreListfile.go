package miniappcloud

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappcloud"
)

// TaobaoMiniappCloudStoreListfile 云存储根据文件名反查地址
// taobao.miniapp.cloud.store.listfile
//
// 云存储中，根据文件名反查地址
func TaobaoMiniappCloudStoreListfile(ctx context.Context, clt *core.SDKClient, req *miniappcloud.TaobaoMiniappCloudStoreListfileAPIRequest, resp *miniappcloud.TaobaoMiniappCloudStoreListfileAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
