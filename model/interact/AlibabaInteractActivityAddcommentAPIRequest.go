package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘评论接口 API请求
alibaba.interact.activity.addcomment

发表评论，并返回楼层
*/
type AlibabaInteractActivityAddcommentAPIRequest struct {
    model.Params
    // 该字段为评论内容
    _content   string
    // 评论feedid
    _feedId   int64
    // 发评论的业务id
    _bizId   string
}

// 初始化AlibabaInteractActivityAddcommentAPIRequest对象
func NewAlibabaInteractActivityAddcommentRequest() *AlibabaInteractActivityAddcommentAPIRequest{
    return &AlibabaInteractActivityAddcommentAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityAddcommentAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.activity.addcomment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityAddcommentAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 该字段为评论内容
func (r *AlibabaInteractActivityAddcommentAPIRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaInteractActivityAddcommentAPIRequest) GetContent() string {
    return r._content
}
// FeedId Setter
// 评论feedid
func (r *AlibabaInteractActivityAddcommentAPIRequest) SetFeedId(_feedId int64) error {
    r._feedId = _feedId
    r.Set("feed_id", _feedId)
    return nil
}

// FeedId Getter
func (r AlibabaInteractActivityAddcommentAPIRequest) GetFeedId() int64 {
    return r._feedId
}
// BizId Setter
// 发评论的业务id
func (r *AlibabaInteractActivityAddcommentAPIRequest) SetBizId(_bizId string) error {
    r._bizId = _bizId
    r.Set("biz_id", _bizId)
    return nil
}

// BizId Getter
func (r AlibabaInteractActivityAddcommentAPIRequest) GetBizId() string {
    return r._bizId
}