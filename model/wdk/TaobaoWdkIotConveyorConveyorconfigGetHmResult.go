package wdk

// TaobaoWdkIotConveyorConveyorconfigGetHMResult 
type TaobaoWdkIotConveyorConveyorconfigGetHMResult struct {
    // 配置信息列表
    Model   *ConveyorBasicConfigDto `json:"model,omitempty" xml:"model,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
