package wdk

// AlibabaWdkOrderGetResult 
type AlibabaWdkOrderGetResult struct {

    // 订单列表
    Orders   []Order `json:"orders,omitempty"`

    // 返回本查询条件下的数据总数
    TotalNumber   int64 `json:"total_number,omitempty"`

    // returnCode
    ReturnCode   string `json:"return_code,omitempty"`

    // returnMsg
    ReturnMsg   string `json:"return_msg,omitempty"`

}