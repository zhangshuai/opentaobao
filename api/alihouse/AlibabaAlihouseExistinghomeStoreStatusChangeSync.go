package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreStatusChangeSync 门店状态变更
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
func AlibabaAlihouseExistinghomeStoreStatusChangeSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreStatusChangeSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
