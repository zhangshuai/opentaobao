package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewFreezeroleAPIRequest 冻结角色 API请求
// alibaba.campus.acl.new.freezerole
//
// 冻结角色
type AlibabaCampusAclNewFreezeroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// NewAlibabaCampusAclNewFreezeroleRequest 初始化AlibabaCampusAclNewFreezeroleAPIRequest对象
func NewAlibabaCampusAclNewFreezeroleRequest() *AlibabaCampusAclNewFreezeroleAPIRequest {
	return &AlibabaCampusAclNewFreezeroleAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusAclNewFreezeroleAPIRequest) Reset() {
	r._workbenchcontext = nil
	r._roleId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewFreezeroleAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.freezerole"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewFreezeroleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusAclNewFreezeroleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewFreezeroleAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabaCampusAclNewFreezeroleAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetRoleId is RoleId Setter
// 角色主键id
func (r *AlibabaCampusAclNewFreezeroleAPIRequest) SetRoleId(_roleId int64) error {
	r._roleId = _roleId
	r.Set("role_id", _roleId)
	return nil
}

// GetRoleId RoleId Getter
func (r AlibabaCampusAclNewFreezeroleAPIRequest) GetRoleId() int64 {
	return r._roleId
}

var poolAlibabaCampusAclNewFreezeroleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusAclNewFreezeroleRequest()
	},
}

// GetAlibabaCampusAclNewFreezeroleRequest 从 sync.Pool 获取 AlibabaCampusAclNewFreezeroleAPIRequest
func GetAlibabaCampusAclNewFreezeroleAPIRequest() *AlibabaCampusAclNewFreezeroleAPIRequest {
	return poolAlibabaCampusAclNewFreezeroleAPIRequest.Get().(*AlibabaCampusAclNewFreezeroleAPIRequest)
}

// ReleaseAlibabaCampusAclNewFreezeroleAPIRequest 将 AlibabaCampusAclNewFreezeroleAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusAclNewFreezeroleAPIRequest(v *AlibabaCampusAclNewFreezeroleAPIRequest) {
	v.Reset()
	poolAlibabaCampusAclNewFreezeroleAPIRequest.Put(v)
}
