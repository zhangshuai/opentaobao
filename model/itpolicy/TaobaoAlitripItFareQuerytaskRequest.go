package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】批量操作结果查询 API请求
taobao.alitrip.it.fare.querytask

批量操作同步返回任务id，后台生成异步任务，通过此接口查询批量操作的执行结果
*/
type TaobaoAlitripItFareQuerytaskRequest struct {
    model.Params
    // json格式的字符串，扩展属性，预留
    extendAttributes   string
    // 任务id
    taskId   int64
}

// 初始化TaobaoAlitripItFareQuerytaskRequest对象
func NewTaobaoAlitripItFareQuerytaskRequest() *TaobaoAlitripItFareQuerytaskRequest{
    return &TaobaoAlitripItFareQuerytaskRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareQuerytaskRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.querytask"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareQuerytaskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareQuerytaskRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItFareQuerytaskRequest) GetExtendAttributes() string {
    return r.extendAttributes
}
// TaskId Setter
// 任务id
func (r *TaobaoAlitripItFareQuerytaskRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("taskId", taskId)
    return nil
}

// TaskId Getter
func (r TaobaoAlitripItFareQuerytaskRequest) GetTaskId() int64 {
    return r.taskId
}
