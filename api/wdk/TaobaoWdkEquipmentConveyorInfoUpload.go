package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentConveyorInfoUpload 五道口仓库悬挂链信息上报
// taobao.wdk.equipment.conveyor.info.upload
//
// 五道口仓库悬挂链信息上传
func TaobaoWdkEquipmentConveyorInfoUpload(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorInfoUploadAPIRequest, resp *wdk.TaobaoWdkEquipmentConveyorInfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
