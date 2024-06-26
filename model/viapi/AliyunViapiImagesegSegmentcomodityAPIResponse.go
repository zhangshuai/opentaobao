package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmentcomodityAPIResponse 商品分割 API返回值
// aliyun.viapi.imageseg.segmentcomodity
//
// 识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmentcomodityAPIResponse struct {
	model.CommonResponse
	AliyunViapiImagesegSegmentcomodityAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmentcomodityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiImagesegSegmentcomodityAPIResponseModel).Reset()
}

// AliyunViapiImagesegSegmentcomodityAPIResponseModel is 商品分割 成功返回结果
type AliyunViapiImagesegSegmentcomodityAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageseg_segmentcomodity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiImagesegSegmentcomodityData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmentcomodityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiImagesegSegmentcomodityAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmentcomodityAPIResponse)
	},
}

// GetAliyunViapiImagesegSegmentcomodityAPIResponse 从 sync.Pool 获取 AliyunViapiImagesegSegmentcomodityAPIResponse
func GetAliyunViapiImagesegSegmentcomodityAPIResponse() *AliyunViapiImagesegSegmentcomodityAPIResponse {
	return poolAliyunViapiImagesegSegmentcomodityAPIResponse.Get().(*AliyunViapiImagesegSegmentcomodityAPIResponse)
}

// ReleaseAliyunViapiImagesegSegmentcomodityAPIResponse 将 AliyunViapiImagesegSegmentcomodityAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiImagesegSegmentcomodityAPIResponse(v *AliyunViapiImagesegSegmentcomodityAPIResponse) {
	v.Reset()
	poolAliyunViapiImagesegSegmentcomodityAPIResponse.Put(v)
}
