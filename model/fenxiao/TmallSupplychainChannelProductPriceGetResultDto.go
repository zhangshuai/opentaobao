package fenxiao

// TmallSupplychainChannelProductPriceGetResultDto 
type TmallSupplychainChannelProductPriceGetResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 询价结果
    Module   *TopProductPriceResult `json:"module,omitempty"`

    // 错误码
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}