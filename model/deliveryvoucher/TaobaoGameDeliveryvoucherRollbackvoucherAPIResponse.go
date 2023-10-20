package deliveryvoucher

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaogamedeliveryvoucherrollbackvoucherAPIResponse 回滚券 API返回值
// taobao.game.deliveryvoucher.rollbackvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaogamedeliveryvoucherrollbackvoucherAPIResponse struct {
	model.CommonResponse
	TaobaogamedeliveryvoucherrollbackvoucherAPIResponseModel
}

// TaobaogamedeliveryvoucherrollbackvoucherAPIResponseModel is 回滚券 成功返回结果
type TaobaogamedeliveryvoucherrollbackvoucherAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_rollbackvoucher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
