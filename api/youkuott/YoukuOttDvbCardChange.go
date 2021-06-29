package youkuott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/youkuott"
)

/* 
dvb ca卡替换 
youku.ott.dvb.card.change

dvb 更换ca卡
*/
func YoukuOttDvbCardChange(clt *core.SDKClient, req *youkuott.YoukuOttDvbCardChangeRequest, session string) (*youkuott.YoukuOttDvbCardChangeAPIResponse, error) {
    var resp youkuott.YoukuOttDvbCardChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}