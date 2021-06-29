package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
对外提供会员注册服务 
alitrip.merchant.galaxy.member.provider.register

星河产品=对外提供注册雅高会员服务
*/
func AlitripMerchantGalaxyMemberProviderRegister(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberProviderRegisterRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberProviderRegisterAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyMemberProviderRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}