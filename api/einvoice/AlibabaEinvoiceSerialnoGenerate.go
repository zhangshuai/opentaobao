package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceSerialnoGenerate 获取统一开票流水号
// alibaba.einvoice.serialno.generate
//
// erp调用开票请求时需要一个开票流水号，此接口就提供了统一的开票流水号，避免了不同系统的冲突
func AlibabaEinvoiceSerialnoGenerate(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceSerialnoGenerateAPIRequest, session string) (*einvoice.AlibabaEinvoiceSerialnoGenerateAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceSerialnoGenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
