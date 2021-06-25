package qimen

// OrderProcessReportRequest 
type OrderProcessReportRequest struct {

    // 订单信息
    Order   *Order `json:"order,omitempty"`

    // 订单处理信息
    Process   *Process `json:"process,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

}