package wdk

// FinanceOrderDetailResponse 
type FinanceOrderDetailResponse struct {

    // 分页信息
    Pagination   *Pagination `json:"pagination,omitempty"`

    // 财务订单信息
    FinanceOrderDetails   []FinanceOrderDetail `json:"finance_order_details,omitempty"`

}