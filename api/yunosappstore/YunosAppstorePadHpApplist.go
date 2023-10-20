package yunosappstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// YunosAppstorePadHpApplist 查询HpPad appList
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
func YunosAppstorePadHpApplist(clt *core.SDKClient, req *yunosappstore.YunosAppstorePadHpApplistAPIRequest, resp *yunosappstore.YunosAppstorePadHpApplistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
