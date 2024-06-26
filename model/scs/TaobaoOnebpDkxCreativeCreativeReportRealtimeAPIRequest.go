package scs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest 获取创意实时报表 API请求
// taobao.onebp.dkx.creative.creative.report.realtime
//
// 获取创意实时报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 创意绑定的查询信息
	_creativeBindQuery *CreativeBindQueryTopDto
}

// NewTaobaoOnebpDkxCreativeCreativeReportRealtimeRequest 初始化TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest对象
func NewTaobaoOnebpDkxCreativeCreativeReportRealtimeRequest() *TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest {
	return &TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) Reset() {
	r._apiServiceContext = nil
	r._creativeBindQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.creative.creative.report.realtime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetCreativeBindQuery is CreativeBindQuery Setter
// 创意绑定的查询信息
func (r *TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) SetCreativeBindQuery(_creativeBindQuery *CreativeBindQueryTopDto) error {
	r._creativeBindQuery = _creativeBindQuery
	r.Set("creative_bind_query", _creativeBindQuery)
	return nil
}

// GetCreativeBindQuery CreativeBindQuery Getter
func (r TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) GetCreativeBindQuery() *CreativeBindQueryTopDto {
	return r._creativeBindQuery
}

var poolTaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOnebpDkxCreativeCreativeReportRealtimeRequest()
	},
}

// GetTaobaoOnebpDkxCreativeCreativeReportRealtimeRequest 从 sync.Pool 获取 TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest
func GetTaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest() *TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest {
	return poolTaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest.Get().(*TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest)
}

// ReleaseTaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest 将 TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest 放入 sync.Pool
func ReleaseTaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest(v *TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest) {
	v.Reset()
	poolTaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest.Put(v)
}
