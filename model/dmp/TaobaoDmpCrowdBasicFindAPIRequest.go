package dmp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDmpCrowdBasicFindAPIRequest DMP_BP版人群列表查询 API请求
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
type TaobaoDmpCrowdBasicFindAPIRequest struct {
	model.Params
	// 请求体
	_apiContext *ApiContextDto
	// 人群查询条件
	_crowdQuery *CrowdQueryDto
	// 分页参数
	_pager *Pager
}

// NewTaobaoDmpCrowdBasicFindRequest 初始化TaobaoDmpCrowdBasicFindAPIRequest对象
func NewTaobaoDmpCrowdBasicFindRequest() *TaobaoDmpCrowdBasicFindAPIRequest {
	return &TaobaoDmpCrowdBasicFindAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDmpCrowdBasicFindAPIRequest) Reset() {
	r._apiContext = nil
	r._crowdQuery = nil
	r._pager = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDmpCrowdBasicFindAPIRequest) GetApiMethodName() string {
	return "taobao.dmp.crowd.basic.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDmpCrowdBasicFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDmpCrowdBasicFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiContext is ApiContext Setter
// 请求体
func (r *TaobaoDmpCrowdBasicFindAPIRequest) SetApiContext(_apiContext *ApiContextDto) error {
	r._apiContext = _apiContext
	r.Set("api_context", _apiContext)
	return nil
}

// GetApiContext ApiContext Getter
func (r TaobaoDmpCrowdBasicFindAPIRequest) GetApiContext() *ApiContextDto {
	return r._apiContext
}

// SetCrowdQuery is CrowdQuery Setter
// 人群查询条件
func (r *TaobaoDmpCrowdBasicFindAPIRequest) SetCrowdQuery(_crowdQuery *CrowdQueryDto) error {
	r._crowdQuery = _crowdQuery
	r.Set("crowd_query", _crowdQuery)
	return nil
}

// GetCrowdQuery CrowdQuery Getter
func (r TaobaoDmpCrowdBasicFindAPIRequest) GetCrowdQuery() *CrowdQueryDto {
	return r._crowdQuery
}

// SetPager is Pager Setter
// 分页参数
func (r *TaobaoDmpCrowdBasicFindAPIRequest) SetPager(_pager *Pager) error {
	r._pager = _pager
	r.Set("pager", _pager)
	return nil
}

// GetPager Pager Getter
func (r TaobaoDmpCrowdBasicFindAPIRequest) GetPager() *Pager {
	return r._pager
}

var poolTaobaoDmpCrowdBasicFindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDmpCrowdBasicFindRequest()
	},
}

// GetTaobaoDmpCrowdBasicFindRequest 从 sync.Pool 获取 TaobaoDmpCrowdBasicFindAPIRequest
func GetTaobaoDmpCrowdBasicFindAPIRequest() *TaobaoDmpCrowdBasicFindAPIRequest {
	return poolTaobaoDmpCrowdBasicFindAPIRequest.Get().(*TaobaoDmpCrowdBasicFindAPIRequest)
}

// ReleaseTaobaoDmpCrowdBasicFindAPIRequest 将 TaobaoDmpCrowdBasicFindAPIRequest 放入 sync.Pool
func ReleaseTaobaoDmpCrowdBasicFindAPIRequest(v *TaobaoDmpCrowdBasicFindAPIRequest) {
	v.Reset()
	poolTaobaoDmpCrowdBasicFindAPIRequest.Put(v)
}
