package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleLevelQuerylevelrule 查询会员等级规则
// alibaba.alsc.crm.rule.level.querylevelrule
//
// 查询会员等级规则
func AlibabaAlscCrmRuleLevelQuerylevelrule(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest, resp *alsc.AlibabaAlscCrmRuleLevelQuerylevelruleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
