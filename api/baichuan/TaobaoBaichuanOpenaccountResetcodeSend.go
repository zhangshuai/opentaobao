package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountResetcodeSend 百川发送找回密码验证码
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
func TaobaoBaichuanOpenaccountResetcodeSend(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountResetcodeSendAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountResetcodeSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
