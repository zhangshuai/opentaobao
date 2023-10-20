package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbownerAdd 民宿房东信息添加
// taobao.xhotel.bnbowner.add
//
// 添加和更新民宿房东的信息
func TaobaoXhotelBnbownerAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbownerAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbownerAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
