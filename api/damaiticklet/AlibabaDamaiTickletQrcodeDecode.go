package damaiticklet

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damaiticklet"
)

// Alibabadamaitickletqrcodedecode 票夹-动态二维码-解码
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
func Alibabadamaitickletqrcodedecode(clt *core.SDKClient, req *damaiticklet.AlibabadamaitickletqrcodedecodeAPIRequest, session string) (*damaiticklet.AlibabadamaitickletqrcodedecodeAPIResponse, error) {
	var resp damaiticklet.AlibabadamaitickletqrcodedecodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
