package ticket

// AlitripTicketRuleUploadResultSet 结构体
type AlitripTicketRuleUploadResultSet struct {
	// 规则维护结果
	FirstResult *TopTicketRuleResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
