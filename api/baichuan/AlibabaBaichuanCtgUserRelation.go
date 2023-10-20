package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanCtgUserRelation 用户
// alibaba.baichuan.ctg.user.relation
//
// 提供给优酷查询道长和淘宝账户的绑定关系
func AlibabaBaichuanCtgUserRelation(clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgUserRelationAPIRequest, resp *baichuan.AlibabaBaichuanCtgUserRelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
