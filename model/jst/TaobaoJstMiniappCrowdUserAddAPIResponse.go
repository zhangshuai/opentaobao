package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序添加用户到指定的活动 API返回值 
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动
*/
type TaobaoJstMiniappCrowdUserAddAPIResponse struct {
    model.CommonResponse
    TaobaoJstMiniappCrowdUserAddAPIResponseModel
}

// 小程序添加用户到指定的活动 成功返回结果
type TaobaoJstMiniappCrowdUserAddAPIResponseModel struct {
    XMLName xml.Name `xml:"jst_miniapp_crowd_user_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 添加成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 请求成功
    RCode   int64 `json:"r_code,omitempty" xml:"r_code,omitempty"`
}