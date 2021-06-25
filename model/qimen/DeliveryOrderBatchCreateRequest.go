package qimen

// DeliveryOrderBatchCreateRequest 
type DeliveryOrderBatchCreateRequest struct {

    // 订单信息
    Orders   []Order `json:"orders,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}