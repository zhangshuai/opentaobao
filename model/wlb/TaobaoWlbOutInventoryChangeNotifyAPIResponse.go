package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlboutinventorychangenotifyAPIResponse 外部库存变化通知（企业物流用户使用） API返回值
// taobao.wlb.out.inventory.change.notify
//
// 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaowlboutinventorychangenotifyAPIResponse struct {
	model.CommonResponse
	TaobaowlboutinventorychangenotifyAPIResponseModel
}

// TaobaowlboutinventorychangenotifyAPIResponseModel is 外部库存变化通知（企业物流用户使用） 成功返回结果
type TaobaowlboutinventorychangenotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_out_inventory_change_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存变化通知成功时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}
