package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest 获取类目过滤条件 API请求
// taobao.universalbp.stdcategory.findcategorycondition
//
// 查询全量类目信息列表，用于进行类目兴趣人群相关定向
type TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// stdCategoryQueryVO
	_stdCategoryQueryVO *StdCategoryQueryVo
}

// NewTaobaouniversalbpstdcategoryfindcategoryconditionRequest 初始化TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest对象
func NewTaobaouniversalbpstdcategoryfindcategoryconditionRequest() *TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest {
	return &TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.stdcategory.findcategorycondition"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetStdCategoryQueryVO is StdCategoryQueryVO Setter
// stdCategoryQueryVO
func (r *TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) SetStdCategoryQueryVO(_stdCategoryQueryVO *StdCategoryQueryVo) error {
	r._stdCategoryQueryVO = _stdCategoryQueryVO
	r.Set("std_category_query_v_o", _stdCategoryQueryVO)
	return nil
}

// GetStdCategoryQueryVO StdCategoryQueryVO Getter
func (r TaobaouniversalbpstdcategoryfindcategoryconditionAPIRequest) GetStdCategoryQueryVO() *StdCategoryQueryVo {
	return r._stdCategoryQueryVO
}