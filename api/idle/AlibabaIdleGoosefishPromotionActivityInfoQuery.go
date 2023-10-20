package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleGoosefishPromotionActivityInfoQuery 闲鱼三方活动参与信息查询
// alibaba.idle.goosefish.promotion.activity.info.query
//
// 闲鱼三方活动参与信息查询
func AlibabaIdleGoosefishPromotionActivityInfoQuery(clt *core.SDKClient, req *idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIRequest, resp *idle.AlibabaIdleGoosefishPromotionActivityInfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
