package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptCreativeGet 获取创意实时报表数据
// taobao.simba.rtrpt.creative.get
//
// 获取创意实时报表数据
func TaobaoSimbaRtrptCreativeGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCreativeGetAPIRequest, resp *simba.TaobaoSimbaRtrptCreativeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
