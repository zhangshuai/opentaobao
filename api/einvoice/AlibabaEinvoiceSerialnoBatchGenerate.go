package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceserialnobatchgenerate 开票流水号批量生成接口
// alibaba.einvoice.serialno.batch.generate
//
// 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
// 优先使用alibaba.einvoice.serial.generate。
func Alibabaeinvoiceserialnobatchgenerate(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceserialnobatchgenerateAPIRequest, session string) (*einvoice.AlibabaeinvoiceserialnobatchgenerateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceserialnobatchgenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
