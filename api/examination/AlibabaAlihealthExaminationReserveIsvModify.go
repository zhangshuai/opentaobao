package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveIsvModify ISV调TOP主动发起改期信息
// alibaba.alihealth.examination.reserve.isv.modify
//
// 体检机构对接_ISV发起体检套餐改期
func AlibabaAlihealthExaminationReserveIsvModify(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveIsvModifyAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveIsvModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
