package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewPageuserroleAPIRequest 分页查询管理员 API请求
// alibaba.campus.acl.new.pageuserrole
//
// 新增用户和角色的关系
type AlibabaCampusAclNewPageuserroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_usersRoleQueryParam *UsersRoleQueryParam
}

// NewAlibabaCampusAclNewPageuserroleRequest 初始化AlibabaCampusAclNewPageuserroleAPIRequest对象
func NewAlibabaCampusAclNewPageuserroleRequest() *AlibabaCampusAclNewPageuserroleAPIRequest {
	return &AlibabaCampusAclNewPageuserroleAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewPageuserroleAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._usersRoleQueryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.pageuserrole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewPageuserroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetUsersRoleQueryParam is UsersRoleQueryParam Setter
// 入参
func (r *AlibabaCampusAclNewPageuserroleAPIRequest) SetUsersRoleQueryParam(_usersRoleQueryParam *UsersRoleQueryParam) error {
	r._usersRoleQueryParam = _usersRoleQueryParam
	r.Set("users_role_query_param", _usersRoleQueryParam)
	return nil
}

// GetUsersRoleQueryParam UsersRoleQueryParam Getter
func (r AlibabaCampusAclNewPageuserroleAPIRequest) GetUsersRoleQueryParam() *UsersRoleQueryParam {
	return r._usersRoleQueryParam
}

var poolAlibabaCampusAclNewPageuserroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewPageuserroleRequest()
	},
}

// GetAlibabaCampusAclNewPageuserroleRequest 从 sync.Pool 获取 AlibabaCampusAclNewPageuserroleAPIRequest
func GetAlibabaCampusAclNewPageuserroleAPIRequest() *AlibabaCampusAclNewPageuserroleAPIRequest {
	return poolAlibabaCampusAclNewPageuserroleAPIRequest.Get().(*AlibabaCampusAclNewPageuserroleAPIRequest)
}

// ReleaseAlibabaCampusAclNewPageuserroleAPIRequest 将 AlibabaCampusAclNewPageuserroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewPageuserroleAPIRequest(v *AlibabaCampusAclNewPageuserroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewPageuserroleAPIRequest.Put(v)
}
