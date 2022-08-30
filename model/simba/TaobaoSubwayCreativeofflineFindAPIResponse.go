package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeofflineFindAPIResponse 获取创意离线多日汇总报表 API返回值
// taobao.subway.creativeoffline.find
//
// 获取创意离线报表
type TaobaoSubwayCreativeofflineFindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCreativeofflineFindAPIResponseModel
}

// TaobaoSubwayCreativeofflineFindAPIResponseModel is 获取创意离线多日汇总报表 成功返回结果
type TaobaoSubwayCreativeofflineFindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_creativeoffline_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result []ReportResultTopDto `json:"result,omitempty" xml:"result>report_result_top_dto,omitempty"`
	// 提示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}