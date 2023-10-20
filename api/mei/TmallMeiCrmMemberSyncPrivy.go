package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// Tmallmeicrmmembersyncprivy 同步推送会员信息(隐私号版本)
// tmall.mei.crm.member.sync.privy
//
// 品牌方通过该api主动推送会员信息。使用场景包括
// 1.用户在线上注册后，线下补充信息后，同步到线上。
// 2.其他情况的主动推送变更。
func Tmallmeicrmmembersyncprivy(clt *core.SDKClient, req *mei.TmallmeicrmmembersyncprivyAPIRequest, session string) (*mei.TmallmeicrmmembersyncprivyAPIResponse, error) {
	var resp mei.TmallmeicrmmembersyncprivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
