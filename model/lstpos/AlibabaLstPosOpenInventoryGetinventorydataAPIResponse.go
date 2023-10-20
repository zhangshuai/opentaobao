package lstpos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopeninventorygetinventorydataAPIResponse 商品库存只读接口(最多20条库存信息) API返回值
// alibaba.lst.pos.open.inventory.getinventorydata
//
// 商品库存只读接口(最多20条库存信息)
type AlibabalstposopeninventorygetinventorydataAPIResponse struct {
	model.CommonResponse
	AlibabalstposopeninventorygetinventorydataAPIResponseModel
}

// AlibabalstposopeninventorygetinventorydataAPIResponseModel is 商品库存只读接口(最多20条库存信息) 成功返回结果
type AlibabalstposopeninventorygetinventorydataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_inventory_getinventorydata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果对象
	Result *AlibabalstposopeninventorygetinventorydataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
