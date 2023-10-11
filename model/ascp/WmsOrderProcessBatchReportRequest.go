package ascp

// WmsOrderProcessBatchReportRequest 结构体
type WmsOrderProcessBatchReportRequest struct {
	// 回告详情列表
	OrderProcessReports []OrderProcessReports `json:"order_process_reports,omitempty" xml:"order_process_reports>order_process_reports,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
