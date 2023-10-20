package koubeimall

// TaobaoKoubeiMallCommonStoreDetailQueryResult 结构体
type TaobaoKoubeiMallCommonStoreDetailQueryResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回门店详情模型
	Data *StoreDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
