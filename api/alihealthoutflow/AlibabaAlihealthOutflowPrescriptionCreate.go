package alihealthoutflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowPrescriptionCreate 处方外流-创建处方
// alibaba.alihealth.outflow.prescription.create
//
// 阿里健康-处方外流-对外提供保存处方功能
func AlibabaAlihealthOutflowPrescriptionCreate(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionCreateAPIRequest, resp *alihealthoutflow.AlibabaAlihealthOutflowPrescriptionCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
