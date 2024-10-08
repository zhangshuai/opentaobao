package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyPutpackageunbind 码拼箱解除父子关系接口
// alibaba.alihealth.drug.kyt.scqy.putpackageunbind
//
// 码拼箱解除父子关系接口
func AlibabaAlihealthDrugKytScqyPutpackageunbind(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyPutpackageunbindAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
