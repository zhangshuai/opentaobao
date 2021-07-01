package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询员工全部权限(包括角色下面的权限) API请求
alibaba.campus.acl.queryallemppermiitem

查询员工全部权限(包括角色下面的权限)
*/
type AlibabaCampusAclQueryallemppermiitemAPIRequest struct {
    model.Params
    // 公司id不填默认SYS_000
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 用户账号
    _accountId   string
    // 每页多少条
    _page   int64
    // 每页记录数
    _pageSize   int64
}

// 初始化AlibabaCampusAclQueryallemppermiitemAPIRequest对象
func NewAlibabaCampusAclQueryallemppermiitemRequest() *AlibabaCampusAclQueryallemppermiitemAPIRequest{
    return &AlibabaCampusAclQueryallemppermiitemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.queryallemppermiitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id不填默认SYS_000
func (r *AlibabaCampusAclQueryallemppermiitemAPIRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclQueryallemppermiitemAPIRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclQueryallemppermiitemAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetCampusId() int64 {
    return r._campusId
}
// AccountId Setter
// 用户账号
func (r *AlibabaCampusAclQueryallemppermiitemAPIRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetAccountId() string {
    return r._accountId
}
// Page Setter
// 每页多少条
func (r *AlibabaCampusAclQueryallemppermiitemAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页记录数
func (r *AlibabaCampusAclQueryallemppermiitemAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaCampusAclQueryallemppermiitemAPIRequest) GetPageSize() int64 {
    return r._pageSize
}