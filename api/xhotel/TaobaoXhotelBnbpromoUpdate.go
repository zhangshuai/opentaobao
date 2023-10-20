package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Taobaoxhotelbnbpromoupdate 民宿营销活动更新
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
func Taobaoxhotelbnbpromoupdate(clt *core.SDKClient, req *xhotel.TaobaoxhotelbnbpromoupdateAPIRequest, session string) (*xhotel.TaobaoxhotelbnbpromoupdateAPIResponse, error) {
	var resp xhotel.TaobaoxhotelbnbpromoupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
