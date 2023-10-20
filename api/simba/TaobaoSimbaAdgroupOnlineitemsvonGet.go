package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupOnlineitemsvonGet 获取用户上架在线销售的全部宝贝
// taobao.simba.adgroup.onlineitemsvon.get
//
// 获取用户上架在线销售的全部宝贝
func TaobaoSimbaAdgroupOnlineitemsvonGet(clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest, resp *simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
