package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaRetailElectronicCertificatePreConfirm 贩卖机开始核销接口
// alibaba.retail.electronic.certificate.pre.confirm
//
// 零售终端贩卖机开始核销接口,返回待领的商品ID
func AlibabaRetailElectronicCertificatePreConfirm(clt *core.SDKClient, req *retail.AlibabaRetailElectronicCertificatePreConfirmAPIRequest, resp *retail.AlibabaRetailElectronicCertificatePreConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
