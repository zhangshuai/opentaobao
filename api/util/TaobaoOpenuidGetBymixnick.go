package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoOpenuidGetBymixnick 通过mixnick转换openuid
// taobao.openuid.get.bymixnick
//
// 通过mixnick转换openuid
func TaobaoOpenuidGetBymixnick(clt *core.SDKClient, req *util.TaobaoOpenuidGetBymixnickAPIRequest, resp *util.TaobaoOpenuidGetBymixnickAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
