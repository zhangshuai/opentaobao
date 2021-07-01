package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
魔盒统计数据查询接口 API请求
yunos.tvpubadmin.data.query

用于华数查询魔盒上的一些用户统计数据
*/
type YunosTvpubadminDataQueryAPIRequest struct {
    model.Params
    // 表名
    _tableName   string
    // 列名
    _columns   string
    // UUID
    _uuid   string
    // 数据类型
    _dataTypeId   int64
    // 日期
    _date   string
    // 页码
    _pageNo   int64
    // 每页个数
    _pageSize   int64
}

// 初始化YunosTvpubadminDataQueryAPIRequest对象
func NewYunosTvpubadminDataQueryRequest() *YunosTvpubadminDataQueryAPIRequest{
    return &YunosTvpubadminDataQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDataQueryAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.data.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDataQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TableName Setter
// 表名
func (r *YunosTvpubadminDataQueryAPIRequest) SetTableName(_tableName string) error {
    r._tableName = _tableName
    r.Set("table_name", _tableName)
    return nil
}

// TableName Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetTableName() string {
    return r._tableName
}
// Columns Setter
// 列名
func (r *YunosTvpubadminDataQueryAPIRequest) SetColumns(_columns string) error {
    r._columns = _columns
    r.Set("columns", _columns)
    return nil
}

// Columns Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetColumns() string {
    return r._columns
}
// Uuid Setter
// UUID
func (r *YunosTvpubadminDataQueryAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetUuid() string {
    return r._uuid
}
// DataTypeId Setter
// 数据类型
func (r *YunosTvpubadminDataQueryAPIRequest) SetDataTypeId(_dataTypeId int64) error {
    r._dataTypeId = _dataTypeId
    r.Set("data_type_id", _dataTypeId)
    return nil
}

// DataTypeId Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetDataTypeId() int64 {
    return r._dataTypeId
}
// Date Setter
// 日期
func (r *YunosTvpubadminDataQueryAPIRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetDate() string {
    return r._date
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminDataQueryAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页个数
func (r *YunosTvpubadminDataQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}