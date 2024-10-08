package drugtrace

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeListCodeAlkali 根据码查询码信息-黄麻碱
// alibaba.alihealth.drug.code.list.code.alkali
//
// 服务描述
// 黄麻碱鉴证核查是基于在两定机构的药品管理（入库、出库或盘点）环节，增加扫码匹配
// 与验证鉴核流程；
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息；
// 若所传的监管码是非最小包装监管码，且存在药品混包的情况，则此接口不支持。这种
// 情况下，需要分多次调用该接口。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
func AlibabaAlihealthDrugCodeListCodeAlkali(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeListCodeAlkaliAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeListCodeAlkaliAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
