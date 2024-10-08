package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelCityCoordinatesBatchUpload 上传信息计算飞猪国际城市
// taobao.xhotel.city.coordinates.batch.upload
//
// 给供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程。目前只适用于国际城市，国内+港澳台暂不支持。
// 非实时计算接口，每次批量上传不少于1条的数据，后端离线计算，请于30分钟后再下载结果。
func TaobaoXhotelCityCoordinatesBatchUpload(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelCityCoordinatesBatchUploadAPIRequest, resp *xhotelitem.TaobaoXhotelCityCoordinatesBatchUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
