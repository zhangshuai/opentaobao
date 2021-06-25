package product

// CainiaoCntecItemChangeMessageResult 
type CainiaoCntecItemChangeMessageResult struct {

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 是否成功接受到请求
    Model   bool `json:"model,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}