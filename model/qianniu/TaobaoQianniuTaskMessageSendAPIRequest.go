package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送任务提醒消息 API请求
taobao.qianniu.task.message.send

如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。
*/
type TaobaoQianniuTaskMessageSendAPIRequest struct {
    model.Params
    // 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
    _taskId   int64
    // 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
    _metadataId   int64
}

// 初始化TaobaoQianniuTaskMessageSendAPIRequest对象
func NewTaobaoQianniuTaskMessageSendRequest() *TaobaoQianniuTaskMessageSendAPIRequest{
    return &TaobaoQianniuTaskMessageSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.message.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskId Setter
// 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
func (r *TaobaoQianniuTaskMessageSendAPIRequest) SetTaskId(_taskId int64) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetTaskId() int64 {
    return r._taskId
}
// MetadataId Setter
// 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
func (r *TaobaoQianniuTaskMessageSendAPIRequest) SetMetadataId(_metadataId int64) error {
    r._metadataId = _metadataId
    r.Set("metadata_id", _metadataId)
    return nil
}

// MetadataId Getter
func (r TaobaoQianniuTaskMessageSendAPIRequest) GetMetadataId() int64 {
    return r._metadataId
}