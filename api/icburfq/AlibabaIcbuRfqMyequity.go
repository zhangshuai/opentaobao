package icburfq

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icburfq"
)

/* 
我的权益 
alibaba.icbu.rfq.myequity

查询供应商权益接口
*/
func AlibabaIcbuRfqMyequity(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqMyequityRequest, session string) (*icburfq.AlibabaIcbuRfqMyequityAPIResponse, error) {
    var resp icburfq.AlibabaIcbuRfqMyequityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}