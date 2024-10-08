package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthMedicalRegistrationSyncnew 阿里健康新挂号数据回传
// alibaba.alihealth.medical.registration.syncnew
//
// 阿里健康新挂号记录回传接口
func AlibabaAlihealthMedicalRegistrationSyncnew(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthMedicalRegistrationSyncnewAPIRequest, resp *alihealth2.AlibabaAlihealthMedicalRegistrationSyncnewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
