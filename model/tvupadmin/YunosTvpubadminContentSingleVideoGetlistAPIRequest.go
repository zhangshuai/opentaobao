package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核获取视频列表 API请求
yunos.tvpubadmin.content.single.video.getlist

牌照方在审核单视频时获取单视频列表接口
*/
type YunosTvpubadminContentSingleVideoGetlistAPIRequest struct {
    model.Params
    // 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
    _extType   int64
    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    _licenseState   int64
    // 单页数量
    _pageSize   int64
    // 查询时间范围，结束时间
    _gmtEnd   string
    // 视频id
    _extVideoStrId   string
    // 查询多个审核状态
    _licenseStateList   []int64
    // 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
    _dateType   int64
    // 主分类
    _category   int64
    // 页码
    _pageNo   int64
    // 查询时间范围，开始时间
    _gmtStart   string
    // 牌照方
    _license   int64
    // 视屏名称
    _videoTitleLike   string
    // 审核优先级，紧急4，高3，中2，低1
    _priority   int64
}

// 初始化YunosTvpubadminContentSingleVideoGetlistAPIRequest对象
func NewYunosTvpubadminContentSingleVideoGetlistRequest() *YunosTvpubadminContentSingleVideoGetlistAPIRequest{
    return &YunosTvpubadminContentSingleVideoGetlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.single.video.getlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtType Setter
// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetExtType(_extType int64) error {
    r._extType = _extType
    r.Set("ext_type", _extType)
    return nil
}

// ExtType Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetExtType() int64 {
    return r._extType
}
// LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetLicenseState(_licenseState int64) error {
    r._licenseState = _licenseState
    r.Set("license_state", _licenseState)
    return nil
}

// LicenseState Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetLicenseState() int64 {
    return r._licenseState
}
// PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// GmtEnd Setter
// 查询时间范围，结束时间
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetGmtEnd(_gmtEnd string) error {
    r._gmtEnd = _gmtEnd
    r.Set("gmt_end", _gmtEnd)
    return nil
}

// GmtEnd Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetGmtEnd() string {
    return r._gmtEnd
}
// ExtVideoStrId Setter
// 视频id
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetExtVideoStrId(_extVideoStrId string) error {
    r._extVideoStrId = _extVideoStrId
    r.Set("ext_video_str_id", _extVideoStrId)
    return nil
}

// ExtVideoStrId Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetExtVideoStrId() string {
    return r._extVideoStrId
}
// LicenseStateList Setter
// 查询多个审核状态
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetLicenseStateList(_licenseStateList []int64) error {
    r._licenseStateList = _licenseStateList
    r.Set("license_state_list", _licenseStateList)
    return nil
}

// LicenseStateList Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetLicenseStateList() []int64 {
    return r._licenseStateList
}
// DateType Setter
// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetDateType(_dateType int64) error {
    r._dateType = _dateType
    r.Set("date_type", _dateType)
    return nil
}

// DateType Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetDateType() int64 {
    return r._dateType
}
// Category Setter
// 主分类
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetCategory(_category int64) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetCategory() int64 {
    return r._category
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// GmtStart Setter
// 查询时间范围，开始时间
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetGmtStart(_gmtStart string) error {
    r._gmtStart = _gmtStart
    r.Set("gmt_start", _gmtStart)
    return nil
}

// GmtStart Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetGmtStart() string {
    return r._gmtStart
}
// License Setter
// 牌照方
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetLicense() int64 {
    return r._license
}
// VideoTitleLike Setter
// 视屏名称
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetVideoTitleLike(_videoTitleLike string) error {
    r._videoTitleLike = _videoTitleLike
    r.Set("video_title_like", _videoTitleLike)
    return nil
}

// VideoTitleLike Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetVideoTitleLike() string {
    return r._videoTitleLike
}
// Priority Setter
// 审核优先级，紧急4，高3，中2，低1
func (r *YunosTvpubadminContentSingleVideoGetlistAPIRequest) SetPriority(_priority int64) error {
    r._priority = _priority
    r.Set("priority", _priority)
    return nil
}

// Priority Getter
func (r YunosTvpubadminContentSingleVideoGetlistAPIRequest) GetPriority() int64 {
    return r._priority
}