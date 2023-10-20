package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Taobaoailabaicloudtopidlistconverter 将淘宝openId或者设备id/用户id转换为openId
// taobao.ailab.aicloud.top.id.list.converter
//
// 将淘宝openId或者设备id/用户id转换为openId
func Taobaoailabaicloudtopidlistconverter(clt *core.SDKClient, req *tmallgenie.TaobaoailabaicloudtopidlistconverterAPIRequest, session string) (*tmallgenie.TaobaoailabaicloudtopidlistconverterAPIResponse, error) {
	var resp tmallgenie.TaobaoailabaicloudtopidlistconverterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}