package aliyun

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliyun"
)

// AccountaliyuncscomcreateAliyunAccountForBid20130701 为bid用户创建账号
// account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01
//
// 给指定的bid创建账号，同时账号打上owner bid的标记
func AccountaliyuncscomcreateAliyunAccountForBid20130701(clt *core.SDKClient, req *aliyun.AccountaliyuncscomcreateAliyunAccountForBid20130701APIRequest, session string) (*aliyun.AccountaliyuncscomcreateAliyunAccountForBid20130701APIResponse, error) {
	var resp aliyun.AccountaliyuncscomcreateAliyunAccountForBid20130701APIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}