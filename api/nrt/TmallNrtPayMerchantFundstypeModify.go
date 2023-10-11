package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtPayMerchantFundstypeModify 修改摊位分账类型
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
func TmallNrtPayMerchantFundstypeModify(clt *core.SDKClient, req *nrt.TmallNrtPayMerchantFundstypeModifyAPIRequest, session string) (*nrt.TmallNrtPayMerchantFundstypeModifyAPIResponse, error) {
	var resp nrt.TmallNrtPayMerchantFundstypeModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
