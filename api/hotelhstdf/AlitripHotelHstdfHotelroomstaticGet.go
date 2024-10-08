package hotelhstdf

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfHotelroomstaticGet 根据类型查询静态字段
// alitrip.hotel.hstdf.hotelroomstatic.get
//
// 根据类型查询分页静态字段
func AlitripHotelHstdfHotelroomstaticGet(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfHotelroomstaticGetAPIRequest, resp *hotelhstdf.AlitripHotelHstdfHotelroomstaticGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
