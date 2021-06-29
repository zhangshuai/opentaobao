package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消轻任务 API请求
taobao.qianniu.task.cancel

由任务发起者调用
*/
type TaobaoQianniuTaskCancelRequest struct {
    model.Params
    // 任务元数据ID
    metaId   int64
    // 任务备注
    memo   string
}

// 初始化TaobaoQianniuTaskCancelRequest对象
func NewTaobaoQianniuTaskCancelRequest() *TaobaoQianniuTaskCancelRequest{
    return &TaobaoQianniuTaskCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskCancelRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MetaId Setter
// 任务元数据ID
func (r *TaobaoQianniuTaskCancelRequest) SetMetaId(metaId int64) error {
    r.metaId = metaId
    r.Set("meta_id", metaId)
    return nil
}

// MetaId Getter
func (r TaobaoQianniuTaskCancelRequest) GetMetaId() int64 {
    return r.metaId
}
// Memo Setter
// 任务备注
func (r *TaobaoQianniuTaskCancelRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r TaobaoQianniuTaskCancelRequest) GetMemo() string {
    return r.memo
}
