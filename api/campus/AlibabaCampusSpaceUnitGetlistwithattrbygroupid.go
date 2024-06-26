package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetlistwithattrbygroupid 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
// alibaba.campus.space.unit.getlistwithattrbygroupid
//
// 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
func AlibabaCampusSpaceUnitGetlistwithattrbygroupid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
