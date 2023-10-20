package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayCreativeofflineFind 获取创意离线多日汇总报表
// taobao.subway.creativeoffline.find
//
// 获取创意离线报表
func TaobaoSubwayCreativeofflineFind(clt *core.SDKClient, req *simba.TaobaoSubwayCreativeofflineFindAPIRequest, resp *simba.TaobaoSubwayCreativeofflineFindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
