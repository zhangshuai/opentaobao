package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsAligenieAlbumsGet 专辑详情
// alibaba.ailabs.aligenie.albums.get
//
// 给予厂商查询专辑下的音频详情
func AlibabaAilabsAligenieAlbumsGet(clt *core.SDKClient, req *iot.AlibabaAilabsAligenieAlbumsGetAPIRequest, resp *iot.AlibabaAilabsAligenieAlbumsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
