package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
记录完成擂台的用户 API请求
tmall.fans.arena.record

记录完成擂台的用户和完成分数
*/
type TmallFansArenaRecordAPIRequest struct {
    model.Params
    // 资金池id
    _cashPoolId   int64
    // 用户得分
    _score   int64
    // mixnick
    _mixNick   string
}

// 初始化TmallFansArenaRecordAPIRequest对象
func NewTmallFansArenaRecordRequest() *TmallFansArenaRecordAPIRequest{
    return &TmallFansArenaRecordAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansArenaRecordAPIRequest) GetApiMethodName() string {
    return "tmall.fans.arena.record"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansArenaRecordAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CashPoolId Setter
// 资金池id
func (r *TmallFansArenaRecordAPIRequest) SetCashPoolId(_cashPoolId int64) error {
    r._cashPoolId = _cashPoolId
    r.Set("cash_pool_id", _cashPoolId)
    return nil
}

// CashPoolId Getter
func (r TmallFansArenaRecordAPIRequest) GetCashPoolId() int64 {
    return r._cashPoolId
}
// Score Setter
// 用户得分
func (r *TmallFansArenaRecordAPIRequest) SetScore(_score int64) error {
    r._score = _score
    r.Set("score", _score)
    return nil
}

// Score Getter
func (r TmallFansArenaRecordAPIRequest) GetScore() int64 {
    return r._score
}
// MixNick Setter
// mixnick
func (r *TmallFansArenaRecordAPIRequest) SetMixNick(_mixNick string) error {
    r._mixNick = _mixNick
    r.Set("mix_nick", _mixNick)
    return nil
}

// MixNick Getter
func (r TmallFansArenaRecordAPIRequest) GetMixNick() string {
    return r._mixNick
}