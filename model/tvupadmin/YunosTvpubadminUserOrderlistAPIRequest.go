package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户订单列表 API请求
yunos.tvpubadmin.user.orderlist

获取用户订单列表
*/
type YunosTvpubadminUserOrderlistAPIRequest struct {
    model.Params
    // 用户ID
    _uid   int64
    // 开始时间
    _createTimeStartStr   string
    // 结束时间
    _createTimeEndStr   string
    // 牌照方
    _license   int64
    // 页码值
    _pageNo   int64
    // 每页行数
    _pageSize   int64
}

// 初始化YunosTvpubadminUserOrderlistAPIRequest对象
func NewYunosTvpubadminUserOrderlistRequest() *YunosTvpubadminUserOrderlistAPIRequest{
    return &YunosTvpubadminUserOrderlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserOrderlistAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.orderlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserOrderlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uid Setter
// 用户ID
func (r *YunosTvpubadminUserOrderlistAPIRequest) SetUid(_uid int64) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r YunosTvpubadminUserOrderlistAPIRequest) GetUid() int64 {
    return r._uid
}
// CreateTimeStartStr Setter
// 开始时间
func (r *YunosTvpubadminUserOrderlistAPIRequest) SetCreateTimeStartStr(_createTimeStartStr string) error {
    r._createTimeStartStr = _createTimeStartStr
    r.Set("create_time_start_str", _createTimeStartStr)
    return nil
}

// CreateTimeStartStr Getter
func (r YunosTvpubadminUserOrderlistAPIRequest) GetCreateTimeStartStr() string {
    return r._createTimeStartStr
}
// CreateTimeEndStr Setter
// 结束时间
func (r *YunosTvpubadminUserOrderlistAPIRequest) SetCreateTimeEndStr(_createTimeEndStr string) error {
    r._createTimeEndStr = _createTimeEndStr
    r.Set("create_time_end_str", _createTimeEndStr)
    return nil
}

// CreateTimeEndStr Getter
func (r YunosTvpubadminUserOrderlistAPIRequest) GetCreateTimeEndStr() string {
    return r._createTimeEndStr
}
// License Setter
// 牌照方
func (r *YunosTvpubadminUserOrderlistAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminUserOrderlistAPIRequest) GetLicense() int64 {
    return r._license
}
// PageNo Setter
// 页码值
func (r *YunosTvpubadminUserOrderlistAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminUserOrderlistAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页行数
func (r *YunosTvpubadminUserOrderlistAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminUserOrderlistAPIRequest) GetPageSize() int64 {
    return r._pageSize
}