package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息ack记录查询 API请求
yunos.service.cmns.coa.message.acks.list

第三方应用开发者调用此接口查询消息ack记录
*/
type YunosServiceCmnsCoaMessageAcksListRequest struct {
    model.Params
    // 消息id
    mid   int64
    // 设备id
    did   int64
    // 分页查询页码
    pageIndex   int64
    // 分页每页数据集数
    pageSize   int64
}

// 初始化YunosServiceCmnsCoaMessageAcksListRequest对象
func NewYunosServiceCmnsCoaMessageAcksListRequest() *YunosServiceCmnsCoaMessageAcksListRequest{
    return &YunosServiceCmnsCoaMessageAcksListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageAcksListRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.acks.list"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageAcksListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mid Setter
// 消息id
func (r *YunosServiceCmnsCoaMessageAcksListRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageAcksListRequest) GetMid() int64 {
    return r.mid
}
// Did Setter
// 设备id
func (r *YunosServiceCmnsCoaMessageAcksListRequest) SetDid(did int64) error {
    r.did = did
    r.Set("did", did)
    return nil
}

// Did Getter
func (r YunosServiceCmnsCoaMessageAcksListRequest) GetDid() int64 {
    return r.did
}
// PageIndex Setter
// 分页查询页码
func (r *YunosServiceCmnsCoaMessageAcksListRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r YunosServiceCmnsCoaMessageAcksListRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// PageSize Setter
// 分页每页数据集数
func (r *YunosServiceCmnsCoaMessageAcksListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosServiceCmnsCoaMessageAcksListRequest) GetPageSize() int64 {
    return r.pageSize
}
