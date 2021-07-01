package flight

// AlitripAgentFlightSellTicketingIssueResultDto 
type AlitripAgentFlightSellTicketingIssueResultDto struct {
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误消息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}