package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryStoreManage 创建或更新仓库
// taobao.inventory.store.manage
//
// 创建商家仓或者更新商家仓信息
func TaobaoInventoryStoreManage(clt *core.SDKClient, req *fenxiao.TaobaoInventoryStoreManageAPIRequest, resp *fenxiao.TaobaoInventoryStoreManageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
