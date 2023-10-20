package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// AlibabaRetailDeviceVendingRegister 贩卖机设备注册
// alibaba.retail.device.vending.register
//
// 贩卖机注册
func AlibabaRetailDeviceVendingRegister(clt *core.SDKClient, req *nrt.AlibabaRetailDeviceVendingRegisterAPIRequest, resp *nrt.AlibabaRetailDeviceVendingRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
