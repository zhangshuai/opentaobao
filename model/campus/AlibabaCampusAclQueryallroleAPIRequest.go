package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询全部角色 API请求
alibaba.campus.acl.queryallrole

查询全部园区
*/
type AlibabaCampusAclQueryallroleAPIRequest struct {
    model.Params
    // 公司id不填统一SYS_000
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 角色名称
    _roleName   string
    // 角色类型
    _roleType   string
    // 角色id
    _roleId   string
}

// 初始化AlibabaCampusAclQueryallroleAPIRequest对象
func NewAlibabaCampusAclQueryallroleRequest() *AlibabaCampusAclQueryallroleAPIRequest{
    return &AlibabaCampusAclQueryallroleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclQueryallroleAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.queryallrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclQueryallroleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id不填统一SYS_000
func (r *AlibabaCampusAclQueryallroleAPIRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclQueryallroleAPIRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclQueryallroleAPIRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclQueryallroleAPIRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclQueryallroleAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclQueryallroleAPIRequest) GetCampusId() int64 {
    return r._campusId
}
// RoleName Setter
// 角色名称
func (r *AlibabaCampusAclQueryallroleAPIRequest) SetRoleName(_roleName string) error {
    r._roleName = _roleName
    r.Set("role_name", _roleName)
    return nil
}

// RoleName Getter
func (r AlibabaCampusAclQueryallroleAPIRequest) GetRoleName() string {
    return r._roleName
}
// RoleType Setter
// 角色类型
func (r *AlibabaCampusAclQueryallroleAPIRequest) SetRoleType(_roleType string) error {
    r._roleType = _roleType
    r.Set("role_type", _roleType)
    return nil
}

// RoleType Getter
func (r AlibabaCampusAclQueryallroleAPIRequest) GetRoleType() string {
    return r._roleType
}
// RoleId Setter
// 角色id
func (r *AlibabaCampusAclQueryallroleAPIRequest) SetRoleId(_roleId string) error {
    r._roleId = _roleId
    r.Set("role_id", _roleId)
    return nil
}

// RoleId Getter
func (r AlibabaCampusAclQueryallroleAPIRequest) GetRoleId() string {
    return r._roleId
}