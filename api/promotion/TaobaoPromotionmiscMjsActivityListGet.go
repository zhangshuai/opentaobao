package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscmjsactivitylistget 查询满就送活动列表
// taobao.promotionmisc.mjs.activity.list.get
//
// 查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
func Taobaopromotionmiscmjsactivitylistget(clt *core.SDKClient, req *promotion.TaobaopromotionmiscmjsactivitylistgetAPIRequest, session string) (*promotion.TaobaopromotionmiscmjsactivitylistgetAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscmjsactivitylistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
