package wdk

// ItemDiscountActivityQueryRequest 结构体
type ItemDiscountActivityQueryRequest struct {
	// erp外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}