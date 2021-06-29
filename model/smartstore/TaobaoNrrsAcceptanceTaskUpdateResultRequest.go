package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新验收任务支付宝结果 API请求
taobao.nrrs.acceptance.task.updateResult

智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。
*/
type TaobaoNrrsAcceptanceTaskUpdateResultRequest struct {
    model.Params
    // 任务ID
    taskId   string
    // 系统自动生成
    alipayResultList   []AlipayCheckResult
}

// 初始化TaobaoNrrsAcceptanceTaskUpdateResultRequest对象
func NewTaobaoNrrsAcceptanceTaskUpdateResultRequest() *TaobaoNrrsAcceptanceTaskUpdateResultRequest{
    return &TaobaoNrrsAcceptanceTaskUpdateResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetApiMethodName() string {
    return "taobao.nrrs.acceptance.task.updateResult"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskId Setter
// 任务ID
func (r *TaobaoNrrsAcceptanceTaskUpdateResultRequest) SetTaskId(taskId string) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetTaskId() string {
    return r.taskId
}
// AlipayResultList Setter
// 系统自动生成
func (r *TaobaoNrrsAcceptanceTaskUpdateResultRequest) SetAlipayResultList(alipayResultList []AlipayCheckResult) error {
    r.alipayResultList = alipayResultList
    r.Set("alipay_result_list", alipayResultList)
    return nil
}

// AlipayResultList Getter
func (r TaobaoNrrsAcceptanceTaskUpdateResultRequest) GetAlipayResultList() []AlipayCheckResult {
    return r.alipayResultList
}
