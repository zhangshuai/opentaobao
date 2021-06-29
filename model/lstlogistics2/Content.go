package lstlogistics2

// Content 
type Content struct {
    // 物流编号，可用来查询物流详情
    LogisticsId   string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
    // 揽收时间
    PickTime   string `json:"pick_time,omitempty" xml:"pick_time,omitempty"`
    // 签收时间
    SignTime   string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
    // 供应商userId
    SupplierUserId   int64 `json:"supplier_user_id,omitempty" xml:"supplier_user_id,omitempty"`
    // 外单主订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 车辆信息
    VehicleInfo   string `json:"vehicle_info,omitempty" xml:"vehicle_info,omitempty"`
    // 司机手机号
    DriverMobile   string `json:"driver_mobile,omitempty" xml:"driver_mobile,omitempty"`
    // 司机
    DriverName   string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
    // 子订单
    SubOrders   []SubOrders `json:"sub_orders,omitempty" xml:"sub_orders>sub_orders,omitempty"`
    // 零售通订单id
    LstOrderId   int64 `json:"lst_order_id,omitempty" xml:"lst_order_id,omitempty"`
    // * 发货单状态 * NEW          ---> 新建 * LOAD_WAIT    ---> 待装车 * LOAD_SUCCESS  ---> 已装车 * SIGN_SUCCESS  ---> 签收、部分签收 * SIGN_FAILED   ---> 拒签 * CANCEL        ---> 取消
    ShipStatus   string `json:"ship_status,omitempty" xml:"ship_status,omitempty"`
    // 出库时间
    OutBoundTime   string `json:"out_bound_time,omitempty" xml:"out_bound_time,omitempty"`
}
