package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有菜单权限 API请求
alibaba.campus.acl.new.checkusermenu

校验用户是否有菜单权限
*/
type AlibabaCampusAclNewCheckusermenuRequest struct {
    model.Params
    // 系统入参
    workbenchcontext   *WorkBenchContext
    // 入参
    param   *CheckUserMenuParam
}

// 初始化AlibabaCampusAclNewCheckusermenuRequest对象
func NewAlibabaCampusAclNewCheckusermenuRequest() *AlibabaCampusAclNewCheckusermenuRequest{
    return &AlibabaCampusAclNewCheckusermenuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewCheckusermenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.checkusermenu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewCheckusermenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewCheckusermenuRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewCheckusermenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// Param Setter
// 入参
func (r *AlibabaCampusAclNewCheckusermenuRequest) SetParam(param *CheckUserMenuParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaCampusAclNewCheckusermenuRequest) GetParam() *CheckUserMenuParam {
    return r.param
}