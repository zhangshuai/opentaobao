package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarXcarSynchronizeCarLinePicsData 爱卡车系图片数据接入
// tmall.car.xcar.synchronize.car.line.pics.data
//
// 爱卡车系图片数据同步天猫汽车
func TmallCarXcarSynchronizeCarLinePicsData(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarXcarSynchronizeCarLinePicsDataAPIRequest, resp *tmallcar.TmallCarXcarSynchronizeCarLinePicsDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
