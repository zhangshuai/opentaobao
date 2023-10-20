package moziacl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoziaclroleremoveAPIRequest 删除角色 API请求
// alibaba.mozi.acl.role.remove
//
// 根据传入的角色code、租户id，删除租户内对应的角色
type AlibabamoziaclroleremoveAPIRequest struct {
	model.Params
	// 删除角色请求对象
	_deleteRolesRequest *DeleteRolesRequest
}

// NewAlibabamoziaclroleremoveRequest 初始化AlibabamoziaclroleremoveAPIRequest对象
func NewAlibabamoziaclroleremoveRequest() *AlibabamoziaclroleremoveAPIRequest {
	return &AlibabamoziaclroleremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoziaclroleremoveAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.acl.role.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoziaclroleremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoziaclroleremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteRolesRequest is DeleteRolesRequest Setter
// 删除角色请求对象
func (r *AlibabamoziaclroleremoveAPIRequest) SetDeleteRolesRequest(_deleteRolesRequest *DeleteRolesRequest) error {
	r._deleteRolesRequest = _deleteRolesRequest
	r.Set("delete_roles_request", _deleteRolesRequest)
	return nil
}

// GetDeleteRolesRequest DeleteRolesRequest Getter
func (r AlibabamoziaclroleremoveAPIRequest) GetDeleteRolesRequest() *DeleteRolesRequest {
	return r._deleteRolesRequest
}
