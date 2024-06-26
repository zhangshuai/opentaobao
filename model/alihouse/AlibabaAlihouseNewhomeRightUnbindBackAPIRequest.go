package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRightUnbindBackAPIRequest 权限回流-解绑 API请求
// alibaba.alihouse.newhome.right.unbind.back
//
// 权限回流-解绑
type AlibabaAlihouseNewhomeRightUnbindBackAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserPermissionInfoDto
}

// NewAlibabaAlihouseNewhomeRightUnbindBackRequest 初始化AlibabaAlihouseNewhomeRightUnbindBackAPIRequest对象
func NewAlibabaAlihouseNewhomeRightUnbindBackRequest() *AlibabaAlihouseNewhomeRightUnbindBackAPIRequest {
	return &AlibabaAlihouseNewhomeRightUnbindBackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) Reset() {
	r._userInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.right.unbind.back"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) SetUserInfo(_userInfo *UserPermissionInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) GetUserInfo() *UserPermissionInfoDto {
	return r._userInfo
}

var poolAlibabaAlihouseNewhomeRightUnbindBackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeRightUnbindBackRequest()
	},
}

// GetAlibabaAlihouseNewhomeRightUnbindBackRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeRightUnbindBackAPIRequest
func GetAlibabaAlihouseNewhomeRightUnbindBackAPIRequest() *AlibabaAlihouseNewhomeRightUnbindBackAPIRequest {
	return poolAlibabaAlihouseNewhomeRightUnbindBackAPIRequest.Get().(*AlibabaAlihouseNewhomeRightUnbindBackAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeRightUnbindBackAPIRequest 将 AlibabaAlihouseNewhomeRightUnbindBackAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRightUnbindBackAPIRequest(v *AlibabaAlihouseNewhomeRightUnbindBackAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRightUnbindBackAPIRequest.Put(v)
}
