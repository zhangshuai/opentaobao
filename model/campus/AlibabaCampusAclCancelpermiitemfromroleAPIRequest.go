package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消角色和权限之间的关系 API请求
alibaba.campus.acl.cancelpermiitemfromrole

取消角色和权限之间的关系
*/
type AlibabaCampusAclCancelpermiitemfromroleAPIRequest struct {
    model.Params
    // 公司ID
    _companyId   int64
    // 系统id
    _systemId   string
    // 园区id
    _campusId   int64
    // 系统自动生成
    _param1   *RoleReq
    // 系统自动生成
    _param2   []PermissionReq
    // 操作人id(不填默认appCode)
    _userId   string
}

// 初始化AlibabaCampusAclCancelpermiitemfromroleAPIRequest对象
func NewAlibabaCampusAclCancelpermiitemfromroleRequest() *AlibabaCampusAclCancelpermiitemfromroleAPIRequest{
    return &AlibabaCampusAclCancelpermiitemfromroleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.cancelpermiitemfromrole"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司ID
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetCompanyId() int64 {
    return r._companyId
}
// SystemId Setter
// 系统id
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetSystemId(_systemId string) error {
    r._systemId = _systemId
    r.Set("system_id", _systemId)
    return nil
}

// SystemId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetSystemId() string {
    return r._systemId
}
// CampusId Setter
// 园区id
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetCampusId() int64 {
    return r._campusId
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetParam1(_param1 *RoleReq) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetParam1() *RoleReq {
    return r._param1
}
// Param2 Setter
// 系统自动生成
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetParam2(_param2 []PermissionReq) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetParam2() []PermissionReq {
    return r._param2
}
// UserId Setter
// 操作人id(不填默认appCode)
func (r *AlibabaCampusAclCancelpermiitemfromroleAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclCancelpermiitemfromroleAPIRequest) GetUserId() string {
    return r._userId
}