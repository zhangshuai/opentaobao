package autonavi

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交通看板-栅格情报获取 APIResponse
alibaba.autonavi.api.trafficboard.image.get

获取指定情报板ID的二进制数据（图片）
*/
type AlibabaAutonaviApiTrafficboardImageGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAutonaviApiTrafficboardImageGetResponse `json:"alibaba_autonavi_api_trafficboard_image_get_response,omitempty"`
}

type AlibabaAutonaviApiTrafficboardImageGetResponse struct {

    // 二进制图片流(png)
    RespResult   []byte `json:"resp_result,omitempty"`

}