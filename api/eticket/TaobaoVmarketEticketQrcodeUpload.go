package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketQrcodeUpload 码商二维码图片上传
// taobao.vmarket.eticket.qrcode.upload
//
// 电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
func TaobaoVmarketEticketQrcodeUpload(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketQrcodeUploadAPIRequest, resp *eticket.TaobaoVmarketEticketQrcodeUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
