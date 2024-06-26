package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaRetailElectronicCertificateConfirm 确认核销接口
// alibaba.retail.electronic.certificate.confirm
//
// 确认核销接口
func AlibabaRetailElectronicCertificateConfirm(clt *core.SDKClient, req *retail.AlibabaRetailElectronicCertificateConfirmAPIRequest, resp *retail.AlibabaRetailElectronicCertificateConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
