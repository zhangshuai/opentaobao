package baodian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户数娱游戏礼包查询 APIResponse
taobao.deg.user.gamegift.query

查询用户数娱礼包列表
*/
type TaobaoDegUserGamegiftQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoDegUserGamegiftQueryResponse `json:"taobao_deg_user_gamegift_query_response,omitempty"`
}

type TaobaoDegUserGamegiftQueryResponse struct {

    // 礼包信息
    Records   []GameGiftRecordDto `json:"records,omitempty"`

}