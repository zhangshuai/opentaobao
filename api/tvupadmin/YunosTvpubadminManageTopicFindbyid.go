package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
根据id获取专题信息 
yunos.tvpubadmin.manage.topic.findbyid

根据id获取专题信息
*/
func YunosTvpubadminManageTopicFindbyid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageTopicFindbyidRequest, session string) (*tvupadmin.YunosTvpubadminManageTopicFindbyidAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageTopicFindbyidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}