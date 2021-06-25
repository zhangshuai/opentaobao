package tmallservice

// TmallServicecenterWorkcardVerifycodeResendResult 
type TmallServicecenterWorkcardVerifycodeResendResult struct {

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

}