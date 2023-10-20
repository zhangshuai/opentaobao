package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsiotbusinessrecipeinsertorupdate 插入和更新食谱
// alibaba.ailabs.iot.business.recipe.insertorupdate
//
// 插入和更新食谱，将isv的食谱添加到云端进行存储
func Alibabaailabsiotbusinessrecipeinsertorupdate(clt *core.SDKClient, req *iot.AlibabaailabsiotbusinessrecipeinsertorupdateAPIRequest, session string) (*iot.AlibabaailabsiotbusinessrecipeinsertorupdateAPIResponse, error) {
	var resp iot.AlibabaailabsiotbusinessrecipeinsertorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
