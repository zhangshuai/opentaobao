package healthnr

// ResponseResult 
type ResponseResult struct {
    // 返回结果
    Result   *LogisticsDetail `json:"result,omitempty" xml:"result,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否成功（true 成功，false 失败）
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}