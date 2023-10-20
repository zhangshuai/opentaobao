package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberDerbyStateSync 德比侧同步卡、券状态接口
// taobao.xhotel.member.derby.state.sync
//
// 德比侧同步卡、券状态接口
func TaobaoXhotelMemberDerbyStateSync(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberDerbyStateSyncAPIRequest, session string) (*xhotelcrm.TaobaoXhotelMemberDerbyStateSyncAPIResponse, error) {
	var resp xhotelcrm.TaobaoXhotelMemberDerbyStateSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
