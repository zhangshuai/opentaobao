package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugcodekythospitalsenddrugmachine 医院发药机
// alibaba.alihealth.drug.code.kyt.hospitalsenddrugmachine
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
// 提供给医院发药机使用
func Alibabaalihealthdrugcodekythospitalsenddrugmachine(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
