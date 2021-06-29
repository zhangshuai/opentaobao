package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片语种识别 API返回值 
alibaba.seaking.imagerecognize

图片语种识别
*/
type AlibabaSeakingImagerecognizeAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingImagerecognizeResponse
}

// 图片语种识别 成功返回结果
type AlibabaSeakingImagerecognizeResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_imagerecognize_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 识别出的图片语种
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}