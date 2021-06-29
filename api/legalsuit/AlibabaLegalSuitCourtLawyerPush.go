package legalsuit

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalsuit"
)

/* 
推荐律师接口 
alibaba.legal.suit.court.lawyer.push

为诉讼系统推荐律师
*/
func AlibabaLegalSuitCourtLawyerPush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtLawyerPushRequest, session string) (*legalsuit.AlibabaLegalSuitCourtLawyerPushAPIResponse, error) {
    var resp legalsuit.AlibabaLegalSuitCourtLawyerPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}