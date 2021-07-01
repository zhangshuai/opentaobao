package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控修改 API请求
yunos.tvpubadmin.content.advert.manageschedule

广告牌照管控修改
*/
type YunosTvpubadminContentAdvertManagescheduleAPIRequest struct {
    model.Params
    // 管理参数
    _req   string
}

// 初始化YunosTvpubadminContentAdvertManagescheduleAPIRequest对象
func NewYunosTvpubadminContentAdvertManagescheduleRequest() *YunosTvpubadminContentAdvertManagescheduleAPIRequest{
    return &YunosTvpubadminContentAdvertManagescheduleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.advert.manageschedule"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Req Setter
// 管理参数
func (r *YunosTvpubadminContentAdvertManagescheduleAPIRequest) SetReq(_req string) error {
    r._req = _req
    r.Set("req", _req)
    return nil
}

// Req Getter
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetReq() string {
    return r._req
}