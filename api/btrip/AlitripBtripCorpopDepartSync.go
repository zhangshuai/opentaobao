package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
外部部门同步 
alitrip.btrip.corpop.depart.sync

同步外部平台部门信息至商旅内部
*/
func AlitripBtripCorpopDepartSync(clt *core.SDKClient, req *btrip.AlitripBtripCorpopDepartSyncRequest, session string) (*btrip.AlitripBtripCorpopDepartSyncAPIResponse, error) {
    var resp btrip.AlitripBtripCorpopDepartSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}