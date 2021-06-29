package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
天猫开新车租后合同下载 
tmall.car.lease.contractdownload

天猫开新车租后合同下载
*/
func TmallCarLeaseContractdownload(clt *core.SDKClient, req *tmallcar.TmallCarLeaseContractdownloadRequest, session string) (*tmallcar.TmallCarLeaseContractdownloadAPIResponse, error) {
    var resp tmallcar.TmallCarLeaseContractdownloadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}