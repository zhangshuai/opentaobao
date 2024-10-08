package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrGetproteminfo 疫苗，获取生产企业的存储和运输温度
// alibaba.alihealth.drug.kyt.dr.getproteminfo
//
// 疫苗，获取生产企业的存储和运输温度
func AlibabaAlihealthDrugKytDrGetproteminfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetproteminfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrGetproteminfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
